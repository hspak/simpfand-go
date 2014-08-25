package main

import (
	"flag"
	"fmt"
)

type config struct {
	incLowTemp  uint16
	incHighTemp uint16
	incMaxTemp  uint16

	incLowLvl    uint16
	incHighLvl   uint16
	incMaxLowLvl uint16

	decLowTemp  uint16
	decHighTemp uint16
	decMaxTemp  uint16

	decLowLvl    uint16
	decHighLvl   uint16
	decMaxLowLvl uint16

	baseLvl uint16

	pollInterval uint16
	maxTemp      uint16
}

func main() {
	flagStart := flag.Bool("start", false, "start simpfand")
	flagStop := flag.Bool("stop", false, "stop simpfand")
	flagVersion := flag.Bool("version", false, "version")
	flagHelp := flag.Bool("help", false, "help")
	flag.Parse()

	if *flagStart == true {
		if moduleExists() {
			fmt.Println("good")
		} else {
			cfg := new(config)
			fanControl(cfg)
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
