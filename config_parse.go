package main

import "os"

func configFileExists() (bool, error) {
	_, err := os.Stat(CONFIG_PATH)
	if err == nil {
		return true, err
	}
	return false, err
}
