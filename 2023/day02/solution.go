// https://adventofcode.com/2023/day/2
// A game (one line of input) consists of one or more rounds (delimited by semicolons) that consist
// of a single handful (delimited by commas)

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

const (
	red   = 12
	green = 13
	blue  = 14
)

// ParseGame takes a string with a game number/round results and returns the game number (if valid, 0 otherwise)
// and the power of the set of minimum cubes for the game to have occurred
func ParseGame(game string) (gameNum, power int) {
	// extract the game identifier and round results
	gameText := game[:strings.Index(game, ":")]
	roundText := game[strings.Index(game, ":")+2:]

	gameNum, err := strconv.Atoi(gameText[strings.Index(gameText, " ")+1:])
	if err != nil {
		panic(err)
	}

	validGame, power := ParseRounds(roundText)
	if !validGame {
		gameNum = 0
	}
	return gameNum, power
}

// ParseRounds takes a string with round results and returns whether the game was possible
// and the power of the set of minimum cubes needed for the game to have occurred
func ParseRounds(rounds string) (valid bool, power int) {
	valid = true // default to a valid round
	var allCounts [3][]int

	roundSplit := strings.Split(rounds, ";")
	// for each round in a game
	for _, round := range roundSplit {
		validHandful, handfulCounts := ParseHandful(round)
		// set round to invalid if any handful is impossible given cube constraints
		if !validHandful {
			valid = false
		}
		// append color count to respective slice if positive
		for i := 0; i < 3; i++ {
			if handfulCounts[i] > 0 {
				allCounts[i] = append(allCounts[i], handfulCounts[i])
			}
		}
	}
	power = CalcPower(allCounts)
	return valid, power
}

// ParseHandful takes a string with handful results and returns whether it was possible
// and the counts for each red, green, and blue cube
func ParseHandful(handful string) (valid bool, counts [3]int) {
	valid = true                             // default to a valid round
	cubeSplit := strings.Split(handful, ",") // split handful into the different colored cubes and their amount pulled

	// for each colored cube
	for _, cubeText := range cubeSplit {
		cubeText = strings.TrimSpace(cubeText)

		// parse cube color and count as an integer
		cubeColor := cubeText[strings.Index(cubeText, " ")+1:]
		cubeCount, err := strconv.Atoi(cubeText[:strings.Index(cubeText, " ")])
		if err != nil {
			panic(err)
		}

		switch cubeColor {
		case "red":
			counts[0] = cubeCount
			if cubeCount > red {
				valid = false
			}
		case "green":
			counts[1] = cubeCount
			if cubeCount > green {
				valid = false
			}
		case "blue":
			counts[2] = cubeCount
			if cubeCount > blue {
				valid = false
			}
		}
	}
	return valid, counts
}

// CalcPower takes the counts of cubes by color and returns the product of their maximum color counts
func CalcPower(counts [3][]int) (power int) {
	power = 1
	for i := 0; i < 3; i++ {
		power *= slices.Max(counts[i])
	}
	return power
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
