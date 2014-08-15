package main

import "github.com/haukurk/occasions/core"
import "code.google.com/p/gcfg"
import "fmt"
import "log"
import "time"
import "os"
import "strconv"

type Configuration struct {
	General struct {
		Port   int
		Listen string
		Ics    string
	}
}

func main() {

	var cfg Configuration

	argsWithoutProg := os.Args[1:]

	// Config file read
	err := gcfg.ReadFileInto(&cfg, "/etc/occasions/occasions.conf") // TODO: Add possibility to define this as an argument.

	if err != nil {
		fmt.Println("Error finding config file")
	}

	if cfg.General.Ics != "" || cfg.General.Port > 0 {
		dates := occasions.ParseDates(cfg.General.Ics)

		if len(argsWithoutProg) > 0 {
			if argsWithoutProg[0] == "rest" {
				occasions.InitRestInterface(strconv.Itoa(cfg.General.Port), dates)
			}
		}

		fmt.Println("[Occasions] Hi! Today is", time.Now().Format("Mon, 01/02/06"))

		err, out, count := occasions.UpcomingDates(dates)

		if err != nil {
			log.Fatal("Error when filtering dates: ", err)
		}

		if count == 0 {
			fmt.Println("[Occasions] Not upcoming occasions found. Sorry :-(")
		} else {
			for _, v := range out {
				fmt.Println("[Occasions] " + v.Greeting + ", " + v.DateStart + " - " + v.Summary)
			}
		}

	} else {
		log.Fatal("vCalendar folder not specified or port not defined in config file.")
	}

}
