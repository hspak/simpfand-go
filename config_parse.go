package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func configFileExists() (bool, error) {
	_, err := os.Stat(CONFIG_PATH)
	if err == nil {
		return true, err
	}
	return false, err
}

func configUpdate(cfg *config, key string, val int) {
	switch key {
	case "POLLING":
		cfg.pollInterval = val
	case "BASE_LVL":
		cfg.baseLvl = val
	case "INC_LOW_TEMP":
		cfg.incLowTemp = val
	case "INC_LOW_LVL":
		cfg.incLowLvl = val
	case "INC_HIGH_TEMP":
		cfg.incHighTemp = val
	case "INC_HIGH_LVL":
		cfg.incHighLvl = val
	case "INC_MAX_TEMP":
		cfg.incMaxTemp = val
	case "INC_MAX_LVL":
		cfg.incMaxLvl = val
	case "DEC_LOW_TEMP":
		cfg.decLowTemp = val
	case "DEC_LOW_LVL":
		cfg.decLowLvl = val
	case "DEC_HIGH_TEMP":
		cfg.decHighTemp = val
	case "DEC_HIGH_LVL":
		cfg.decHighLvl = val
	case "DEC_MAX_TEMP":
		cfg.decMaxTemp = val
	case "DEC_MAX_LVL":
		cfg.decMaxLvl = val
	}
}

func configParse(cfg *config) (bool, error) {
	const KEY = 0
	const VAL = 1
	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read := scanner.Text()
		if len(read) > 0 && read[0] != '#' {
			pair := strings.Split(read, "=")
			trimmedKey := strings.TrimSpace(pair[KEY])
			trimmedVal := strings.TrimSpace(pair[VAL])
			val, _ := strconv.Atoi(trimmedVal)
			configUpdate(cfg, trimmedKey, val)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error")
	}

	return true, err
}
