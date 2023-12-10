package day2

import (
	"strconv"
	"strings"
)

const (
	MAXIMUM_RED_CUBES   = 12
	MAXIMUM_GREEN_CUBES = 13
	MAXIMUM_BLUE_CUBES  = 14
)

type Day2 struct {
}

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Cubes []Cubes
}

func parseGameString(gameString string) (*Game, error) {
	var game Game

	for _, phase := range strings.Split(gameString, ";") {
		cube := Cubes{}
		for _, cubes := range strings.Split(phase, ",") {
			cubes = strings.TrimSpace(cubes)
			phase := strings.Split(cubes, " ")
			count, err := strconv.Atoi(phase[0])
			if err != nil {
				return nil, err
			}
			color := phase[1]

			switch color {
			case "red":
				cube.Red = count
			case "green":
				cube.Green = count
			case "blue":
				cube.Blue = count
			}

		}
		game.Cubes = append(game.Cubes, cube)
	}
	return &game, nil
}

func prepareGame(line string) (string, int, error) {
	gameFullString := strings.Split(line, ":")
	gameNumber, err := strconv.Atoi(strings.Split(gameFullString[0], " ")[1])
	if err != nil {
		return "", 0, err
	}
	gameString := strings.TrimSpace(gameFullString[1])
	return gameString, gameNumber, nil
}

func (d Day2) Part1(input []string) (string, error) {
	var sum int
	for _, line := range input {
		gameString, gameNumber, err := prepareGame(line)
		if err != nil {
			return "", err
		}

		game, err := parseGameString(gameString)
		if err != nil {
			return "", err
		}

		var possibleGame bool = true
		for _, cube := range game.Cubes {
			if cube.Red > MAXIMUM_RED_CUBES || cube.Green > MAXIMUM_GREEN_CUBES || cube.Blue > MAXIMUM_BLUE_CUBES {
				possibleGame = false
			}
		}

		if possibleGame {
			sum += gameNumber
		}
	}

	return strconv.Itoa(sum), nil
}

func findMiximumCubesPerGame(game *Game) *Cubes {
	var minimumCubes Cubes
	for _, cube := range game.Cubes {
		if cube.Red > minimumCubes.Red {
			minimumCubes.Red = cube.Red
		}
		if cube.Green > minimumCubes.Green {
			minimumCubes.Green = cube.Green
		}
		if cube.Blue > minimumCubes.Blue {
			minimumCubes.Blue = cube.Blue
		}
	}
	return &minimumCubes
}

func (d Day2) Part2(input []string) (string, error) {
	var sum int
	for _, line := range input {
		gameString, _, err := prepareGame(line)
		if err != nil {
			return "", err
		}

		game, err := parseGameString(gameString)
		if err != nil {
			return "", err
		}

		minimumCubes := findMiximumCubesPerGame(game)
		product := minimumCubes.Red * minimumCubes.Green * minimumCubes.Blue
		sum += product

	}

	return strconv.Itoa(sum), nil
}
