package main

import (
	"flag"
	"fmt"
)

type config struct {
	incLowTemp  int
	incHighTemp int
	incMaxTemp  int

	incLowLvl  int
	incHighLvl int
	incMaxLvl  int

	decLowTemp  int
	decHighTemp int
	decMaxTemp  int

	decLowLvl  int
	decHighLvl int
	decMaxLvl  int

	baseLvl      int
	pollInterval int
}

func createConfig() *config {
	cfg := new(config)

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

func main() {
	flagStart := flag.Bool("start", false, "start simpfand")
	flagStop := flag.Bool("stop", false, "stop simpfand")
	flagVersion := flag.Bool("version", false, "version")
	flagHelp := flag.Bool("help", false, "help")
	flag.Parse()

	if *flagStart == true {
		if moduleExists() {
			cfg := createConfig()
			configParse(cfg)
			fanControl(cfg)
			fmt.Println("good")
		} else {
			fmt.Println("bad")
		}
	} else if *flagStop {
		fmt.Println("stop")
	} else if *flagVersion {
		fmt.Println("version")
	} else if *flagHelp || true {
		fmt.Println("help")
	}
}
