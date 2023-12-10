package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day4 struct {
}

type Card struct {
	Number      int
	WinningHand []int
	MyHand      []int
	Power       int
}

func parseLine(line string) (*Card, error) {
	winningHand := []int{}
	myHand := []int{}
	line = strings.Replace(line, "   ", " ", -1)
	line = strings.Replace(line, "  ", " ", -1)
	cardString := strings.Split(line, ":")

	cardNumber, err := strconv.Atoi(strings.Split(cardString[0], " ")[1])
	if err != nil {
		return nil, fmt.Errorf("cardNumber, error parsing %s: %w", cardString[0], err)
	}
	hands := strings.Split(cardString[1], "|")
	winningHandString := strings.TrimSpace(hands[0])
	myHandString := strings.TrimSpace(hands[1])
	for _, n := range strings.Split(winningHandString, " ") {
		n = strings.TrimSpace(n)
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("winningHand, error parsing %s: %w", n, err)
		}
		winningHand = append(winningHand, number)
	}

	for _, n := range strings.Split(myHandString, " ") {
		n = strings.TrimSpace(n)
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("myHand, error parsing %s: %w", n, err)
		}
		myHand = append(myHand, number)
	}

	return &Card{
		Number:      cardNumber,
		WinningHand: winningHand,
		MyHand:      myHand,
		Power:       1,
	}, nil
}

func getSameCardCount(cards *Card) int {
	var count int
	for _, wc := range cards.WinningHand {
		for _, mc := range cards.MyHand {
			if wc == mc {
				count++
				break
			}
		}
	}

	return count
}

func getCardWinnings(cards *Card) int {
	count := getSameCardCount(cards)

	return int(math.Pow(2, float64(count-1)))
}

func (d Day4) Part1(input []string) (string, error) {
	var total int
	for _, line := range input {
		card, err := parseLine(line)
		if err != nil {
			return "", err
		}
		total += getCardWinnings(card)
	}

	return strconv.Itoa(total), nil
}

func (d Day4) Part2(input []string) (string, error) {
	var scratchcardsCount int

	var cards []*Card

	for _, line := range input {
		card, err := parseLine(line)
		if err != nil {
			return "", err
		}

		cards = append(cards, card)
	}

	for i := 0; i < len(cards); i++ {
		count := getSameCardCount(cards[i])
		for j := 0; j < count; j++ {
			cards[i+j+1].Power += cards[i].Power
		}
	}

	for _, card := range cards {
		scratchcardsCount += card.Power
	}
	return strconv.Itoa(scratchcardsCount), nil

}
