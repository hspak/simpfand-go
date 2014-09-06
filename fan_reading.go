package main

import (
	"bufio"
	"fmt"
	"log"
	"log/syslog"
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

func getTemp() uint16 {
	var file *os.File
	validPath := false
	for _, path := range tempPaths {
		var err error
		file, err = os.Open(path)
		if err != nil {
			continue
		}
		validPath = true
	}

	if !validPath {
		log.Fatal("Error: could not find temperature readings")
	}

	var read string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read = scanner.Text()
	}
	temp, _ := strconv.Atoi(read)
	return uint16(temp / 1000)
}

func getFanLevel(cfg *config, currTemp uint16, prevTemp uint16, prevLvl uint16) uint16 {
	currLvl := prevLvl
	temp_diff := currTemp - prevTemp

	if temp_diff > 0 {
		if currTemp <= cfg.incLowTemp {
			currLvl = cfg.baseLvl
		} else if currTemp <= cfg.incHighTemp {
			currLvl = cfg.incLowLvl
		} else if currTemp <= cfg.incMaxTemp {
			currLvl = cfg.incHighLvl
		} else {
			currLvl = cfg.incMaxLvl
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

func setFanLevel(lvl uint16) {
	lvlStr := "level " + strconv.Itoa(int(lvl))
	file, err := os.OpenFile(FAN_PATH, os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Error: could not open the fan path.")
	}
	defer file.Close()

	_, err = file.Write([]byte(lvlStr))
	if err != nil {
		log.Fatal("Error: could not write to fan level.")
	}
}

func fanControl(cfg *config, mainLogger *syslog.Writer) {
	currTemp := getTemp()
	prevTemp := currTemp
	prevLvl := uint16(0)
	currLvl := getFanLevel(cfg, currTemp, prevTemp, prevLvl)
	for {
		if prevLvl != currLvl {
			setFanLevel(currLvl)
			mainLogger.Debug(fmt.Sprintf("Fan level changed: %d -> %d (%d -> %d)",
				prevLvl, currLvl, prevTemp, currTemp))
		} else {
			mainLogger.Debug(fmt.Sprintf("Fan level remained: %d (%d -> %d)", currLvl, prevTemp, currTemp))
		}

		time.Sleep(time.Second * time.Duration(cfg.pollInterval))
		prevTemp = currTemp
		currTemp = getTemp()
		prevLvl = currLvl
		currLvl = getFanLevel(cfg, currTemp, prevTemp, prevLvl)
	}
}
