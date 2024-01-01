// https://adventofcode.com/2023/day/6

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// StrToIntSlice converts a string containing integers delimited by a space into an integer slice
func StrToIntSlice(str string) []int {
	stringSlice := strings.Split(str, " ")
	intSlice := []int{}
	for i := range stringSlice {
		if stringSlice[i] != "" {
			var err error
			integer, err := strconv.Atoi(stringSlice[i])
			if err != nil {
				panic(err)
			}
			intSlice = append(intSlice, integer)
		}
	}
	return intSlice
}

func Distance(timeLimit, timeHeld int) int {
	timeMoving := timeLimit - timeHeld
	boatSpeed := timeHeld
	return timeMoving * boatSpeed
}

func main() {
	filename := "input.txt"
	// filename := "input_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var timeLimit []int
	var distanceRecord []int
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		ints := strings.TrimSpace(input[strings.Index(input, ":")+1:])
		if strings.Index(input, "Time") == 0 {
			timeLimit = StrToIntSlice(ints)
		} else {
			distanceRecord = StrToIntSlice(ints)
		}
	}
	// fmt.Printf("Time: %v\nDistance: %v", timeLimit, distanceRecord)

	waysToWin := make([]int, len(timeLimit))
	for i := range timeLimit {
		for hold := 0; hold <= timeLimit[i]; hold++ {
			distance := Distance(timeLimit[i], hold)
			if distance > distanceRecord[i] {
				waysToWin[i]++
			}
		}
	}

	partOne := 1
	for i := range waysToWin {
		partOne *= waysToWin[i]
	}
	fmt.Printf("Product of ways to win (part one): %v\n", partOne)
}
