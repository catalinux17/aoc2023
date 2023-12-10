package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(dayNumber int, inputFile string) []string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	var filePath string
	filePath = fmt.Sprintf("%s/pkg/days/%d/%s", currentDir, dayNumber, inputFile)
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
