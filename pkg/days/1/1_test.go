package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1_Part1(t *testing.T) {
	tests := map[string]struct {
		input   []string
		want    string
		wantErr bool
	}{
		"simple": {
			input: []string{
				"1htuaentu2",
			},
			want:    "12",
			wantErr: false,
		},
		"simple2": {
			input: []string{
				"1htuaentu2",
				"1htuaentu2",
			},
			want:    "24",
			wantErr: false,
		},
		"sample1": {
			input: []string{
				"treb7uchet",
			},
			want:    "77",
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day1{}
			got, err := d.Part1(tt.input)
			if err != nil {
				assert.True(t, tt.wantErr)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_convertStringsToNumber(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"simple1": {
			input: "one",
			want:  "1",
		},
		"simple2": {
			input: "onetwo",
			want:  "12",
		},
		"medium1": {
			input: "onetwothree",
			want:  "123",
		},
		"medium2": {
			input: "tharueightwo",
			want:  "8",
		},
		"complex1": {
			input: "1tharueightwo5three",
			want:  "1853",
		},
		"complex2": {
			input: "teauhnninetuhanth8ttua",
			want:  "98",
		},
		"sample1": {
			input: "1abc2",
			want:  "12",
		},
		"sample2": {
			input: "pqr3stu8vwx",
			want:  "38",
		},
		"sample3": {
			input: "a1b2c3d4e5f",
			want:  "12345",
		},
		"sample4": {
			input: "treb7uchet",
			want:  "7",
		},
		"sample5": {
			input: "9dlvndqbddghpxc",
			want:  "9",
		},
		"sample6": {
			input: "rtkrbtthree8sixfoureight6",
			want:  "386486",
		},
		"sample7": {
			input: "hxcfone64ninesevenbgsnrqppqmnnineeightwof",
			want:  "1649798",
		},
		"sample8": {
			input: "threebjpbtpzgx5mnthreensixoneightz",
			want:  "35361",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := convertStringsToNumber(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDay1_Part2(t *testing.T) {
	tests := map[string]struct {
		input   []string
		want    string
		wantErr bool
	}{
		"simple": {
			input: []string{
				"1htuaentu2",
			},
			want:    "12",
			wantErr: false,
		},
		"simple2": {
			input: []string{
				"1htuaentu2",
				"1htuaentu2",
			},
			want:    "24",
			wantErr: false,
		},
		"medium1": {
			input: []string{
				"1htuaeone",
				"1htuaentu2",
			},
			want:    "23",
			wantErr: false,
		},
		"complex1": {
			input: []string{
				"1htuaeone",
				"1htuaentu2",
				"teauhnninetuhanth8ttua",
			},
			want:    "121",
			wantErr: false,
		},
		"sample1": {
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			want:    "281",
			wantErr: false,
		},
		"sample2": {
			input: []string{
				"7one1tnqxfvhmjvjzfive",
				"sevenmcjs3lmlmxmcgptwobjggfive6four",
				"seven8five3",
				"5sfknxsn5sevenfour446",
				"bxc5two67seven2",
				"jcsfivefive89seven85",
				"nine296",
			},
			want:    "481",
			wantErr: false,
		},
		"sample3": {
			input: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			want:    "142",
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day1{}
			got, err := d.Part2(tt.input)
			if err != nil {
				assert.True(t, tt.wantErr)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
