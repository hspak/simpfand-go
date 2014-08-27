package main

import (
	"bufio"
	"fmt"
	"log"
	"log/syslog"
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

func configUpdate(cfg *config, key string, val uint16) {
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

func configParse(cfg *config, mainLogger *syslog.Writer) bool {
	const KEY = 0
	const VAL = 1
	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	newConf := false
	for scanner.Scan() {
		read := scanner.Text()
		if len(read) > 0 && read[0] != '#' {
			newConf = true
			pair := strings.Split(read, "=")
			trimmedKey := strings.TrimSpace(pair[KEY])
			trimmedVal := strings.TrimSpace(pair[VAL])
			val, _ := strconv.Atoi(trimmedVal)
			configUpdate(cfg, trimmedKey, uint16(val))
			mainLogger.Info(fmt.Sprintf("Custom setting found %s: %s", trimmedKey, trimmedVal))
		}
	}

	if !newConf {
		mainLogger.Info("No custom settings found, using defaults")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not parse config file")
	}

	return true
}
