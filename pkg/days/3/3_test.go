package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getNumbers(t *testing.T) {
	tests := map[string]struct {
		lines []string
		want  []Number
	}{
		"simple": {
			lines: []string{
				"10..20",
				"30..40",
			},
			want: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      0,
					StartCol: 4,
					EndCol:   5,
					Number:   20,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   30,
				},
				{
					Row:      1,
					StartCol: 4,
					EndCol:   5,
					Number:   40,
				},
			},
		},
		"medium": {
			lines: []string{
				"10..20..30.",
				"40..50..60.",
				"70..80..90.",
			},
			want: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      0,
					StartCol: 4,
					EndCol:   5,
					Number:   20,
				},
				{
					Row:      0,
					StartCol: 8,
					EndCol:   9,
					Number:   30,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   40,
				},
				{
					Row:      1,
					StartCol: 4,

					EndCol: 5,
					Number: 50,
				},
				{
					Row:      1,
					StartCol: 8,
					EndCol:   9,
					Number:   60,
				},
				{
					Row:      2,
					StartCol: 0,
					EndCol:   1,
					Number:   70,
				},
				{
					Row:      2,
					StartCol: 4,
					EndCol:   5,
					Number:   80,
				},
				{
					Row:      2,
					StartCol: 8,
					EndCol:   9,
					Number:   90,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getNumbers(tc.lines)
			assert.Equal(t, tc.want, actual)
		})
	}
}

func Test_getAdjacentNumbersToSymbols(t *testing.T) {
	tests := map[string]struct {
		lines   []string
		numbers []Number
		want    []Number
	}{
		"simple": {
			lines: []string{
				"10+.20",
				"30..40",
			},
			numbers: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      0,
					StartCol: 4,
					EndCol:   5,
					Number:   20,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   30,
				},
				{
					Row:      1,
					StartCol: 4,
					EndCol:   5,
					Number:   40,
				},
			},
			want: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   30,
				},
			},
		},
		"withSymbols": {
			lines: []string{
				"10+.20",
				"30*-40",
			},
			numbers: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      0,
					StartCol: 4,
					EndCol:   5,
					Number:   20,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   30,
				},
				{
					Row:      1,
					StartCol: 4,
					EndCol:   5,
					Number:   40,
				},
			},
			want: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      0,
					StartCol: 4,
					EndCol:   5,
					Number:   20,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   30,
				},
				{
					Row:      1,
					StartCol: 4,
					EndCol:   5,
					Number:   40,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getAdjacentNumbersToSymbols(tc.lines, tc.numbers)
			assert.Equal(t, tc.want, actual)
		})
	}
}

func Test_getProduct(t *testing.T) {
	tests := map[string]struct {
		lines   []string
		numbers []Number
		want    int
	}{
		"simple": {
			lines: []string{
				"10*.20",
				"30..40",
			},
			numbers: []Number{
				{
					Row:      0,
					StartCol: 0,
					EndCol:   1,
					Number:   10,
				},
				{
					Row:      0,
					StartCol: 4,
					EndCol:   5,
					Number:   20,
				},
				{
					Row:      1,
					StartCol: 0,
					EndCol:   1,
					Number:   30,
				},
				{
					Row:      1,
					StartCol: 4,
					EndCol:   5,
					Number:   40,
				},
			},
			want: 300,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getProduct(tc.lines, tc.numbers)
			assert.Equal(t, tc.want, actual)
		})
	}
}
