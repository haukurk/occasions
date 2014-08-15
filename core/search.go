package occasions

import (
	"log"
	"strings"
	"time"
)

func UpcomingDates(dates []ODate) (error_function error, out []ODate, count int) {
	//out := make([]string, 0)
	layout := "Mon, 01/02/06"

	now := time.Now()
	year, month, day := now.Date()

	for _, d := range dates {
		timestart, err := time.Parse("20060102", strings.Trim(d.DateStart, " "))
		timeend, err := time.Parse("20060102", strings.Trim(d.DateEnd, " "))
		error_function = err
		if err != nil {
			log.Fatal("Error when parsing time from parsed vCals: ", err)
			break
		}
		curr_year, curr_month, curr_day := timestart.Date()
		today := ((curr_day == day) && (curr_month == month) && (curr_year == year))
		if timestart.After(now) || today {
			switch {
			case today:
				d.DateStart = timestart.Format(layout)
				d.DateEnd = timeend.Format(layout)
				d.Greeting = "Today"
				out = append(out, d)
				count++
			case timestart.Before(now.AddDate(0, 0, 1)):
				d.DateStart = timestart.Format(layout)
				d.DateEnd = timeend.Format(layout)
				d.Greeting = "Tomorrow"
				out = append(out, d)
				count++
			case timestart.Before(now.AddDate(0, 0, 7)):
				d.DateStart = timestart.Format(layout)
				d.DateEnd = timeend.Format(layout)
				d.Greeting = "Next Week"
				out = append(out, d)
				count++
			}
		}
	}
	return
}
