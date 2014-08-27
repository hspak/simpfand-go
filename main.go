package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
)

var VERSION string
var tempPaths [TEMP_PATH_CNT]string

type config struct {
	incLowTemp  uint16
	incHighTemp uint16
	incMaxTemp  uint16

	incLowLvl  uint16
	incHighLvl uint16
	incMaxLvl  uint16

	decLowTemp  uint16
	decHighTemp uint16
	decMaxTemp  uint16

	decLowLvl  uint16
	decHighLvl uint16
	decMaxLvl  uint16

	maxLvl       uint16
	baseLvl      uint16
	pollInterval uint16
}

func createConfig() *config {
	cfg := new(config)

	// defaults
	cfg.incLowLvl = 2
	cfg.incHighLvl = 4
	cfg.incMaxLvl = 6

	cfg.incLowTemp = 55
	cfg.incHighTemp = 65
	cfg.incMaxTemp = 82

	cfg.decLowLvl = 4
	cfg.decHighLvl = 5
	cfg.decMaxLvl = 6

	cfg.decLowTemp = 50
	cfg.decHighTemp = 60
	cfg.decMaxTemp = 77

	cfg.baseLvl = 1
	cfg.pollInterval = 10

	return cfg
}

func showHelp() {
	fmt.Printf("Usage: simpfand-go <action>\n\n" +

		"  Actions:\n" +
		"    --start       start simpfand-go\n" +
		"    --version     display version\n" +
		"    --help        display help\n\n" +

		"NOTE: running --start manually is not recommended!\n")
}

func showVersion() {
	fmt.Println("simpfand-go version:", VERSION)
}

func main() {
	// will take the executable name with empty tag
	mainLogger, err := syslog.New(syslog.LOG_ERR, "")
	if err != nil {
		log.Fatal("Error: could not start syslog")
	}

	flagStart := flag.Bool("start", false, "start simpfand")
	flagVersion := flag.Bool("version", false, "version")
	flagHelp := flag.Bool("help", false, "help")
	flag.Parse()

	if *flagStart == true {
		if moduleExists() {
			cfg := createConfig()
			if !configParse(cfg, mainLogger) {
				mainLogger.Err("Error: could not open config file, using defaults")
			}
			fanControl(cfg, mainLogger)
		} else {
			os.Exit(1)
		}
	} else if *flagVersion {
		showVersion()
	} else if *flagHelp || true {
		showHelp()
	}
}

func init() {
	tempPaths[0] = TEMP_PATH_1
	tempPaths[1] = TEMP_PATH_2
}
