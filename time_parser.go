package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file := openFile("test_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		match := timeRegExp.FindStringSubmatch(line)
		result := make(map[string]string)
		hour, min, meridiem := getResultsFromTimeParse(match, result)
		DayMins := getMinutesOfDay(hour, min, meridiem)

		fmt.Println("Time: ", line)
		fmt.Println("Minutes past midnight: ", strconv.Itoa(DayMins))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	closeFile(file)
}

var timeRegExp = regexp.MustCompile(`(?P<hours>\d\d?)(:(?P<minutes>\d\d))?(?P<meridiem>am|pm)?`)

func openFile(filename string) (file *os.File) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func closeFile(fp *os.File) {
	fp.Close()
}

func getMinutesOfDay(hour int, minutes int, meridiem string) (DayMins int) {
	if meridiem == "am" && hour == 12 {
		DayMins = minutes
		return
	}

	if meridiem == "pm" && hour == 12 {
		DayMins = hour*60 + minutes
		return
	}

	if meridiem == "pm" {
		hour = hour + 12
	}

	DayMins = hour*60 + minutes

	return

}

func getResultsFromTimeParse(match []string, result map[string]string) (hour int, min int, meridiem string) {

	for i, name := range timeRegExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	hour, err1 := strconv.Atoi(result["hours"])

	if err1 != nil {
		return
	}

	min = 0
	var err2 error

	if result["minutes"] != "" {
		min, err2 = strconv.Atoi(result["minutes"])
		if err2 != nil {
			return
		}
	}

	meridiem = result["meridiem"]

	return

}
