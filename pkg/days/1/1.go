package day1

import (
	"fmt"
	"unicode"
)

type Day1 struct {
}

func trimLetters(line string) string {
	var newLine string
	for _, char := range line {
		if unicode.IsDigit(char) {
			newLine += string(char)
		}
	}
	return newLine
}

func (d Day1) Part1(input []string) (string, error) {
	var sum int
	for _, line := range input {
		newLine := trimLetters(line)
		firstNumber := int(newLine[0] - '0')
		lastNumber := int(newLine[len(newLine)-1] - '0')
		sum += firstNumber*10 + lastNumber
	}

	return fmt.Sprint(sum), nil
}

var numbersMap = map[string]string{
	// "zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func convertStringsToNumber(line string) string {
	var newLine string
	var found bool
	var i int
	for i = 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			newLine += string(line[i])
			continue
		}

		for key, value := range numbersMap {
			if i+len(key) > len(line) {
				continue
			}
			if line[i:i+len(key)] == key {
				newLine += value
				found = true
				i += len(key) - 2
				break
			}
		}

		if found {
			found = false
			continue
		}
	}
	return newLine

}

func (d Day1) Part2(input []string) (string, error) {
	var sum int

	for _, line := range input {
		newLine := convertStringsToNumber(line)
		firstNumber := int(newLine[0] - '0')
		lastNumber := int(newLine[len(newLine)-1] - '0')
		sum += firstNumber*10 + lastNumber
	}

	return fmt.Sprint(sum), nil
}
