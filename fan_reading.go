package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func moduleExists() bool {
	file, err := os.Open(FAN_PATH)
	if err != nil {
		fmt.Println("Error: thinkpad_acpi option fan_control is not enabled.")
		return false
	}
	defer file.Close()
	return true
}

func getTemp() int {
	file, err := os.Open(TEMP_PATH)
	if err != nil {
		fmt.Println("Error: could not read temperatures.")
		os.Exit(1)
	}

	var read string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read = scanner.Text()
	}
	temp, _ := strconv.Atoi(read)
	return temp / 1000
}

func getFanLevel(cfg *config, curr_temp int, prev_temp int) int {
	newLvl := cfg.baseLvl
	temp_diff := curr_temp - prev_temp

	if temp_diff > 0 {
		if curr_temp <= cfg.incLowTemp {
			newLvl = cfg.incLowLvl
		} else if curr_temp <= cfg.incHighTemp {
			newLvl = cfg.incHighLvl
		} else if curr_temp <= cfg.incMaxTemp {
			newLvl = cfg.incMaxLvl
		} else {
			newLvl = cfg.maxLvl
		}
	} else if temp_diff < 0 {
		if curr_temp > cfg.decMaxTemp {
			newLvl = cfg.decMaxLvl
		} else if curr_temp > cfg.decHighTemp {
			newLvl = cfg.decHighLvl
		} else if curr_temp > cfg.decLowTemp {
			newLvl = cfg.decLowLvl
		} else {
			newLvl = cfg.baseLvl
		}
	}
	return newLvl
}

func setFanLevel(lvl int) {
	lvlStr := "level " + string(lvl)
	file, err := os.Open(FAN_PATH)
	if err != nil {
		fmt.Println("Error: could not open the fan path.")
		os.Exit(1)
	}
	defer file.Close()
	file.Write([]byte(lvlStr))
}

func fanControl(cfg *config) {
	curr_temp := getTemp()
	prev_temp := curr_temp

	old_lvl := 0
	new_lvl := getFanLevel(cfg, curr_temp, prev_temp)
	for {
		fmt.Println("curr_temp", curr_temp)
		if old_lvl != new_lvl {
			setFanLevel(new_lvl)
		}

		time.Sleep(time.Second * time.Duration(cfg.pollInterval))
		prev_temp = curr_temp
		curr_temp = getTemp()
	}
}
