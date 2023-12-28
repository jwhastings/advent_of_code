package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const red int = 12
const green int = 13
const blue int = 14

func ParseGame(game string) (gameNum, power int) {
	gameText := game[:strings.Index(game, ":")]
	roundText := game[strings.Index(game, ":")+2:]

	gameNum, err := strconv.Atoi(gameText[strings.Index(gameText, " ")+1:])
	if err != nil {
		panic(err)
	}

	validGame, power := ParseRound(roundText)
	if !validGame {
		gameNum = 0
	}

	return gameNum, power
}

func ParseRound(rounds string) (valid bool, power int) {
	valid = true // default to a valid round
	var redCounts []int
	var greenCounts []int
	var blueCounts []int

	roundSplit := strings.Split(rounds, ";")
	// for each round
	for _, round := range roundSplit {
		cubeSplit := strings.Split(round, ",")
		// for each colored cube
		for _, cubeText := range cubeSplit {
			cubeText = strings.TrimSpace(cubeText)

			cubeCount, err := strconv.Atoi(cubeText[:strings.Index(cubeText, " ")])
			if err != nil {
				panic(err)
			}
			cubeColor := cubeText[strings.Index(cubeText, " ")+1:]

			switch cubeColor {
			case "red":
				redCounts = append(redCounts, cubeCount)
				if cubeCount > red {
					valid = false
				}
			case "green":
				greenCounts = append(greenCounts, cubeCount)
				if cubeCount > green {
					valid = false
				}
			case "blue":
				blueCounts = append(blueCounts, cubeCount)
				if cubeCount > blue {
					valid = false
				}
			}
		}
	}

	maxRed := slices.Max(redCounts)
	maxGreen := slices.Max(greenCounts)
	maxBlue := slices.Max(blueCounts)
	power = maxRed * maxGreen * maxBlue

	return valid, power
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sumPartOne := 0
	sumPartTwo := 0
	for scanner.Scan() {
		stringValue := scanner.Text()
		partOne, partTwo := ParseGame(stringValue)
		sumPartOne += partOne
		sumPartTwo += partTwo
	}

	fmt.Printf("Sum of valid game IDs: %v\n", sumPartOne)
	fmt.Printf("Sum of power of sets: %v\n", sumPartTwo)

	file.Close()
}
