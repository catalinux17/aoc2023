package day3

import (
	"fmt"
	"unicode"
)

type Day3 struct {
}

var symbols = []string{"+", "-", "*", "/"}

type Number struct {
	Row      int
	StartCol int
	EndCol   int
	Number   int
}

func getNumbers(lines []string) []Number {
	var numbers []Number

	for i, line := range lines {
		var number int
		var numberFound bool
		for j, c := range line {
			if unicode.IsDigit(c) {
				numberFound = true
				number = number*10 + int(c-'0')
				if j == len(line)-1 {
					numbers = append(numbers, Number{
						Row:      i,
						StartCol: j - len(fmt.Sprintf("%d", number)) + 1,
						EndCol:   j,
						Number:   number,
					})
				}
				continue
			}
			if numberFound {
				numbers = append(numbers, Number{
					Row:      i,
					StartCol: j - len(fmt.Sprintf("%d", number)),
					EndCol:   j - 1,
					Number:   number,
				})
				number = 0
				numberFound = false
			}
		}
	}

	return numbers
}

func getAdjacentNumbersToSymbols(lines []string, numbers []Number) []Number {
	var goodNumbers []Number
	var i, j int
	for _, n := range numbers {
		var found bool
		for i = 0; i < len(lines) && !found; i++ {
			for j = 0; j < len(lines[i]) && !found; j++ {
				if lines[i][j] == '.' || unicode.IsDigit(rune(lines[i][j])) {
					continue
				}
				if (i == n.Row-1 || i == n.Row || i == n.Row+1) && (j >= n.StartCol-1 && j <= n.EndCol+1) {
					goodNumbers = append(goodNumbers, n)
					found = true
				}
			}
		}
	}

	return goodNumbers
}

func getProduct(lines []string, numbers []Number) int {
	var i, j int
	var productSum int
	for _, n := range numbers {
		var found bool
		for i = 0; i < len(lines) && !found; i++ {
			for j = 0; j < len(lines[i]) && !found; j++ {
				if lines[i][j] != '*' {
					continue
				}
				if (i == n.Row-1 || i == n.Row || i == n.Row+1) && (j >= n.StartCol-1 && j <= n.EndCol+1) {
					for _, n2 := range numbers {
						if n2.Row == n.Row && n2.StartCol == n.StartCol {
							continue
						}
						if (i == n2.Row-1 || i == n2.Row || i == n2.Row+1) && (j >= n2.StartCol-1 && j <= n2.EndCol+1) {
							productSum += n.Number * n2.Number
						}
						found = true
					}
				}
			}
		}
	}

	return productSum / 2
}

func (d Day3) Part1(input []string) (string, error) {
	numebers := getNumbers(input)
	goodNumbers := getAdjacentNumbersToSymbols(input, numebers)
	var sum int
	for _, n := range goodNumbers {
		sum += n.Number
	}
	return fmt.Sprintf("%d", sum), nil
}

func (d Day3) Part2(input []string) (string, error) {
	numbers := getNumbers(input)
	product := getProduct(input, numbers)
	return fmt.Sprintf("%d", product), nil
}
