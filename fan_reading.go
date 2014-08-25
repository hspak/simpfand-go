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

func getFanLevel() string {
	file, err := os.Open(FAN_PATH)
	if err != nil {
		fmt.Println("Error: could not open the fan path.")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	lvl, _ := scanner.ReadBytes('\n')
	return string(lvl)
}

func setFanLevel(cfg *config, curr_temp int, prev_temp int) {

	currLvl := getFanLevel()
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

	newLvlStr := string(newLvl)
	if temp_diff != 0 && newLvlStr != currLvl {
		file, err := os.Open(FAN_PATH)
		if err != nil {
			fmt.Println("Error: could not open the fan path.")
			os.Exit(1)
		}
		defer file.Close()
		file.Write([]byte(newLvlStr))
	}
}

func fanControl(cfg *config) {
	curr_temp := getTemp()
	prev_temp := curr_temp

	// curr_lvl :=
	for {
		fmt.Println("curr_temp", curr_temp)
		setFanLevel(cfg, curr_temp, prev_temp)

		time.Sleep(time.Second * time.Duration(cfg.pollInterval))
		prev_temp = curr_temp
		curr_temp = getTemp()
	}
}
