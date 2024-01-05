// https://adventofcode.com/2023/day/8

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

// SameValues
func SameValues(s []int) bool {
	var firstVal int
	for i, val := range s {
		if i == 0 {
			firstVal = val
		} else if firstVal != val {
			return false
		}
	}
	return true
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

	var partOne, partTwo int
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		history := StrToIntSlice(input)
		var histories [][]int
		histories = append(histories, history)
		i := 0
		var lastVal, firstVal int
		for !SameValues(histories[i]) {
			nextHistory := make([]int, len(histories[i])-1)
			for j := range nextHistory {
				nextHistory[j] = histories[i][j+1] - histories[i][j]
			}
			histories = append(histories, nextHistory)
			i++

			if SameValues(nextHistory) {
				lastVal = 0
				firstVal = histories[i][0]
				for k := i; k >= 0; k-- {
					if k < i {
						firstVal = histories[k][0] - firstVal
					}
					lastVal += histories[k][len(histories[k])-1]
				}
			}
		}

		partOne += lastVal
		partTwo += firstVal
	}
	fmt.Printf("Sum of extrapolated values (part one): %v\n", partOne)
	fmt.Printf("Sum of backward-extrapolated values (part two): %v\n", partTwo)
}
