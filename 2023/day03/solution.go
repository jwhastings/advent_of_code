package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	width      int    = 140
	digits     string = "0123456789"
	notsymbols string = ".0123456789"
)

func AppendNewToSlice(slice []string, char string) []string {
	if IsSymbol(char) {
		if len(slice) == 0 {
			return append(slice, char)
		}
		for _, ele := range slice {
			if ele == char {
				return slice
			}
		}
		return append(slice, char)
	}
	return slice
}

func IsSymbol(char string) bool {
	return !(strings.ContainsAny(char, notsymbols))
}

func CheckRange(a, b int, line string) bool {
	for i := a; i < b; i++ {
		char := string(line[i])
		if IsSymbol(char) {
			return true
		}
	}
	return false

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var schematic []string
	var symbols []string
	for scanner.Scan() {
		// get string from input file, line by line, append to slice
		stringValue := scanner.Text()
		schematic = append(schematic, stringValue)
		// iterate over line to identify new symbols
		for _, char := range stringValue {
			symbols = AppendNewToSlice(symbols, string(char))
		}
	}
	file.Close()

	var partNumbers []int
	var adjacentSymbol bool
	for row := 0; row < len(schematic); row++ {
		// for row := 0; row < 3; row++ {
		line := schematic[row]
		re := regexp.MustCompile("[0-9]+")
		allDigits := re.FindAllIndex([]byte(line), -1)
		var numStr string
		for _, ab := range allDigits {
			adjacentSymbol = false
			a := ab[0] // index for start of number
			b := ab[1] // index of end of number + 1
			numStr = line[a:b]
			// check horizontals
			charToLeft := (a > 0)
			charToRight := (b < width)
			charAbove := row > 0
			charBelow := row < len(schematic)-1

			if charToLeft && CheckRange(a-1, a, line) {
				adjacentSymbol = true
			}
			if charToRight && CheckRange(b, b+1, line) {
				adjacentSymbol = true
			}
			// check above
			if charAbove {
				lineAbove := schematic[row-1]

				if charToLeft && charToRight && CheckRange(a-1, b+1, lineAbove) {
					adjacentSymbol = true
				} else if charToLeft && CheckRange(a-1, b, lineAbove) {
					adjacentSymbol = true
				} else if charToRight && CheckRange(a, b+1, lineAbove) {
					adjacentSymbol = true
				} else if CheckRange(a, b, lineAbove) {
					adjacentSymbol = true
				}
			}
			// check below
			if charBelow {
				lineBelow := schematic[row+1]

				if charToLeft && charToRight && CheckRange(a-1, b+1, lineBelow) {
					adjacentSymbol = true
				} else if charToLeft && CheckRange(a-1, b, lineBelow) {
					adjacentSymbol = true
				} else if charToRight && CheckRange(a, b+1, lineBelow) {
					adjacentSymbol = true
				} else if CheckRange(a, b, lineBelow) {
					adjacentSymbol = true
				}
			}
			if adjacentSymbol {
				partInt, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				partNumbers = append(partNumbers, partInt)
			}
		}
	}

	sumPartOne := 0
	for _, num := range partNumbers {
		sumPartOne += num
	}
	fmt.Printf("Sum of numbers with adjacent symbol: %v\n", sumPartOne)
}
