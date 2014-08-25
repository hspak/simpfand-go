package main

import (
	"fmt"
	"os"
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

// func readFanLevel

func fanControl(cfg *config) {
	cfg.baseLvl = 1
	fmt.Println(*cfg)
}
