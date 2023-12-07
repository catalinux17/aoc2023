package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(dayNumber int, test bool) []string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	var filePath string
	if test {
		filePath = fmt.Sprintf("%s/pkg/days/%d/test_input", currentDir, dayNumber)
	} else {
		filePath = fmt.Sprintf("%s/pkg/days/%d/input", currentDir, dayNumber)
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
