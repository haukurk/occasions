package main

import "github.com/haukurk/occasions/utils"
import "code.google.com/p/gcfg"
import "fmt"
import "log"

type Configuration struct {
	General struct {
		Port   int
		Listen string
		Ics    string
	}
}

func main() {

	var cfg Configuration

	// Config file read
	err := gcfg.ReadFileInto(&cfg, "/etc/occasions/occasions.conf")

	if err != nil {
		fmt.Println("Error finding config file")
	} else {
		//fmt.Println("No error")
	}

	if cfg.General.Ics != "" {
		dates := utils.ParseDates(cfg.General.Ics)

		for _, d := range dates {
			fmt.Println(d.Summary, d.DateStart)
		}

	} else {
		log.Fatal("vCalendar folder not specified in config file.")
	}

}
