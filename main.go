package main

import (
	"aoc2023/pkg/days"
	day1 "aoc2023/pkg/days/1"
	day2 "aoc2023/pkg/days/2"
	day3 "aoc2023/pkg/days/3"
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
	case 2:
		day = day2.Day2{}
	case 3:
		day = day3.Day3{}
	}

	input := utils.ReadInput(dayNumber, "input")

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
