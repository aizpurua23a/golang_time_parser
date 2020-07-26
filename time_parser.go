package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

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

var timeRegExp = regexp.MustCompile(`(?P<hours>\d\d?)(:(?P<minutes>\d\d))?(?P<meridiem>am|pm)?`)

func main() {

	file := openFile("test_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := timeRegExp.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range timeRegExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		fmt.Printf("hours: %s\nminutes: %s\nmeridiem: %s\n", result["hours"], result["minutes"], result["meridiem"])

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	closeFile(file)
}
