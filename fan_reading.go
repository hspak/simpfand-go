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

func getFanLevel(cfg *config, currTemp int, prevTemp int) int {
	currLvl := cfg.baseLvl
	temp_diff := currTemp - prevTemp

	if temp_diff > 0 {
		if currTemp <= cfg.incLowTemp {
			currLvl = cfg.incLowLvl
		} else if currTemp <= cfg.incHighTemp {
			currLvl = cfg.incHighLvl
		} else if currTemp <= cfg.incMaxTemp {
			currLvl = cfg.incMaxLvl
		} else {
			currLvl = cfg.maxLvl
		}
	} else if temp_diff < 0 {
		if currTemp > cfg.decMaxTemp {
			currLvl = cfg.decMaxLvl
		} else if currTemp > cfg.decHighTemp {
			currLvl = cfg.decHighLvl
		} else if currTemp > cfg.decLowTemp {
			currLvl = cfg.decLowLvl
		} else {
			currLvl = cfg.baseLvl
		}
	}
	return currLvl
}

func setFanLevel(lvl int) {
	lvlStr := "level " + string(lvl)
	file, err := os.Open(FAN_PATH)
	if err != nil {
		fmt.Println("Error: could not open the fan path.")
		os.Exit(1)
	}
	defer file.Close()

	write, err := file.Write([]byte(lvlStr))
	if err != nil {
		fmt.Println("Error: could not set fan level", write)
	}
}

func fanControl(cfg *config) {
	currTemp := getTemp()
	prevTemp := currTemp
	prevLvl := 0
	currLvl := getFanLevel(cfg, currTemp, prevTemp)
	for {
		if prevLvl != currLvl {
			setFanLevel(currLvl)
		}

		time.Sleep(time.Second * time.Duration(cfg.pollInterval))
		prevTemp = currTemp
		currTemp = getTemp()
		prevLvl = currLvl
		currLvl = getFanLevel(cfg, currTemp, prevTemp)
	}
}
