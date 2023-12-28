package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// getFirstChar returns the index and value of the first digit occuring in a string
func getFirstChar(s string) (firstIndex int, firstDigit string) {
	firstIndex = strings.IndexAny(s, "0123456789") // find the index of the first digit
	firstDigit = string(s[firstIndex])             // extract the digit at given index
	return firstIndex, firstDigit
}

// getLastChar returns the index and value of the last digit occuring in a string
func getLastChar(s string) (lastIndex int, lastDigit string) {
	lastIndex = strings.LastIndexAny(s, "0123456789") // find the index of the last digit
	lastDigit = string(s[lastIndex])                  // extract the digit at given index
	return lastIndex, lastDigit
}

// combineDigits takes two digits as strings and combines them into a two digit number
func combineDigits(first, last string) int {
	twoDigitNumber, err := strconv.Atoi(first + last)
	if err != nil {
		panic(err)
	}
	return twoDigitNumber
}

// partOne calculates the two digit number for a given string
func partOne(s string) (result int) {
	_, firstDigit := getFirstChar(s)
	_, lastDigit := getLastChar(s)
	result = combineDigits(firstDigit, lastDigit)
	return result
}

func convertCharToNumber(s string) (convertedString string) {
	words := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	numbers := []string{
		"ze0ro",
		"on1e",
		"tw2o",
		"thr3ee",
		"fo4ur",
		"fi5ve",
		"si6x",
		"sev7en",
		"ei8ht",
		"n9ne",
	}

	convertedString = s

	for i, word := range words {
		convertedString = strings.ReplaceAll(convertedString, word, numbers[i])
	}
	return convertedString
}

func partTwo(s string) (result int, newStr string) {
	newStr = convertCharToNumber(s)
	result = partOne(newStr)
	return result, newStr
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var sumPartOne, sumPartTwo int
	test := 0
	for scanner.Scan() {
		// get string in current line
		stringValue := scanner.Text()

		// convert digits to an integer and add to total
		resultOne := partOne(stringValue)
		resultTwo, newStringValue := partTwo(stringValue)
		if false {
			fmt.Println(newStringValue)
		}
		sumPartOne += resultOne
		sumPartTwo += resultTwo
		test += 1
		// if test > 5 {
		// 	break
		// }
	}

	file.Close()

	fmt.Printf("part 1 answer: %d\npart 2 answer: %d", sumPartOne, sumPartTwo)
}
