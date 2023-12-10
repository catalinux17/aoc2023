package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseLine(t *testing.T) {
	tests := map[string]struct {
		line string
		want *Card
	}{
		"simple": {
			line: "Card 1: 1 2 3 4 5 | 6 7 8 9 10",
			want: &Card{
				Number:      1,
				WinningHand: []int{1, 2, 3, 4, 5},
				MyHand:      []int{6, 7, 8, 9, 10},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseLine(tc.line)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDay4_Part1(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"simple": {
			input: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			want: "13",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day4{}
			got, err := d.Part1(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}

}
