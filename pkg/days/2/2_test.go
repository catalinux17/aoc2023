package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prepareGame(t *testing.T) {
	tests := map[string]struct {
		input      string
		gameString string
		gameNumber int
	}{
		"simple": {
			input:      "Game 1: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
			gameString: "5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
			gameNumber: 1,
		},
		"simple 2": {
			input:      "Game 2: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
			gameString: "5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
			gameNumber: 2,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, got1, err := prepareGame(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.gameString, got)
			assert.Equal(t, tt.gameNumber, got1)
		})
	}
}

func Test_parseGameString(t *testing.T) {
	tests := map[string]struct {
		input string
		want  *Game
	}{
		"simple": {
			input: "5 red, 5 green, 5 blue",
			want: &Game{
				Cubes: []Cubes{
					{
						Red:   5,
						Green: 5,
						Blue:  5,
					},
				},
			},
		},
		"medium": {
			input: "5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue",
			want: &Game{
				Cubes: []Cubes{
					{
						Red:   5,
						Green: 5,
						Blue:  5,
					},
					{
						Red:   4,
						Green: 4,
						Blue:  4,
					},
					{
						Red:   3,
						Green: 3,
						Blue:  3,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseGameString(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDay2_Part1(t *testing.T) {

	tests := map[string]struct {
		input []string
		want  string
	}{
		"simple": {
			input: []string{
				"Game 1: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
				"Game 2: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
				"Game 3: 5 red, 5 green, 20 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
			},
			want: "3",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day2{}
			got, err := d.Part1(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDay2_Part2(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"simple": {
			input: []string{
				"Game 1: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue; 2 red, 2 green, 2 blue; 1 red, 1 green, 1 blue",
			},
			want: "125",
		},
		"medium": {
			input: []string{
				"Game 1: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue",
				"Game 2: 5 red, 5 green, 5 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue",
				"Game 3: 5 red, 5 green, 20 blue; 4 red, 4 green, 4 blue; 3 red, 3 green, 3 blue",
			},
			want: "750",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day2{}
			got, err := d.Part2(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
