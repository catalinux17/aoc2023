package main

import (
	"aoc2023/pkg/days"
	day1 "aoc2023/pkg/days/1"
	"aoc2023/pkg/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func SelectDay() {
	var day days.Day

	dayString := os.Getenv("DAY")
	if dayString == "" {
		fmt.Println("No day selected")
		return
	}
	dayNumber, err := strconv.Atoi(dayString)
	if err != nil {
		log.Fatal(err)
	}

	switch dayNumber {
	case 1:
		day = day1.Day1{}
	}

	input := utils.ReadInput(dayNumber, false)

	fmt.Println("Day number: ", dayNumber)
	output, err := day.Part1(input)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("part 1: ", output)

	output, err = day.Part2(input)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("part 2: ", output)

}
func main() {
	SelectDay()
}
