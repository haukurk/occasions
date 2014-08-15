package main

import "github.com/haukurk/occasions/utils"
import "github.com/haukurk/occasions/modules"
import "code.google.com/p/gcfg"
import "fmt"
import "log"
import "time"

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
	err := gcfg.ReadFileInto(&cfg, "/etc/occasions/occasions.conf") // TODO: Add possibility to define this as an argument.

	if err != nil {
		fmt.Println("Error finding config file")
	}

	if cfg.General.Ics != "" {
		dates := utils.ParseDates(cfg.General.Ics)

		occasions := 0
		now := time.Now()
		year, month, day := now.Date()

		fmt.Println("[Occasions] Hi, today is ", now)
		fmt.Println("[Occasions] Checking for upcoming occasions..")

		for _, d := range dates {
			timestart, err := time.Parse("20060102", d.DateStart)
			if err != nil {
				log.Fatal("Error when parsing time from parsed vCals: ", err)
				break
			}
			curr_year, curr_month, curr_day := timestart.Date()
			// Ignore old events.
			today := ((curr_day == day) && (curr_month == month) && (curr_year == year))
			if timestart.After(now) || today {
				switch {
				case today:
					fmt.Println("Today!", timestart, d.Summary)
					occasions++
				case timestart.Before(now.AddDate(0, 0, 1)):
					fmt.Println("Tomorrow!", timestart, d.Summary)
					occasions++
				case timestart.Before(now.AddDate(0, 0, 7)):
					fmt.Println("Next Week!", timestart, d.Summary)
					occasions++
				}
			}
		}
		if occasions == 0 {
			fmt.Println("[Occasions] Not upcoming occasions found. Sorry :-(")
		}

	} else {
		log.Fatal("vCalendar folder not specified in config file.")
	}

}
