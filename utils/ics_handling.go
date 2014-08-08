package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type ODate struct {
	Summary   string
	DateStart string
	DateEnd   string
}

// Exported Functions

func ParseDates(folder string) []ODate {

	files, _ := ioutil.ReadDir(folder)
	errors := make([]string, 0)
	dates := make([]ODate, 0)

	for _, f := range files {
		lock := false
		lines, err := readLines("/etc/occasions/ics/" + f.Name())
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		if !strings.Contains(lines[0], "BEGIN:VCALENDAR") && !strings.Contains(lines[len(lines)-1], "END:VCALENDAR") {
			message := fmt.Sprintf("Possibly a malformed ics file (%s)", f.Name())
			errors = append(errors, message)
		}

		tmpOdate := ODate{}

		for i, line := range lines {
			if lock == true {
				if strings.Contains(line, "BEGIN:VEVENT") {
					message := fmt.Sprintf("Malformed ics file (%s) line: %d", f.Name(), i)
					//log.Fatalf(message)
					errors = append(errors, message)
				} else if strings.Contains(line, "END:VEVENT") {
					lock = false
					// Add to Dates Slice
					dates = append(dates, tmpOdate)
				} else {
					re_summary := regexp.MustCompile(`SUMMARY:(.*)`)
					re_start := regexp.MustCompile(`DTSTART;VALUE=DATE:(.*)`)
					re_end := regexp.MustCompile(`DTEND;VALUE=DATE:(.*)`)
					if re_summary.MatchString(line) {
						tmpOdate.Summary = re_summary.FindStringSubmatch(line)[1]
					} else if re_start.MatchString(line) {
						tmpOdate.DateStart = re_start.FindStringSubmatch(line)[1]
					} else if re_end.MatchString(line) {
						tmpOdate.DateEnd = re_end.FindStringSubmatch(line)[1]
					}
				}
			} else {
				if strings.Contains(line, "BEGIN:VEVENT") {
					lock = true
				} else if strings.Contains(line, "END:VEVENT") {
					message := fmt.Sprintf("Malformed ics file (%s) line: %d", f.Name(), i)
					//log.Fatalf(message)
					errors = append(errors, message)
				}
			}
		}
	}

	if len(errors) == 0 {
		//fmt.Println(dates)
		//log.Print("All vCalendar files parsed successfully")
	} else {
		log.Printf("Errors occured when parsing your vCalendar files in %s", folder)
		log.Fatal(errors)
	}

	return dates
}

// Internal Functions

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
