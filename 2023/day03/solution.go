package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func CheckLeft(a, b int, line string) bool {
	return IsSymbol(string(line[a-1]))
}

func CheckRight(a, b int, line string) bool {
	return IsSymbol(string(line[b]))
}

func CheckAboveBelow(a, b int, lineAboveBelow string) bool {
	for i := a; i < b; i++ {
		if IsSymbol(string(lineAboveBelow[i])) {
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

	fmt.Println(strings.Repeat("=", 140))
	// for row := 0; row < len(schematic); row++ {
	for row := 0; row < 3; row++ {
		line := schematic[row]
		fmt.Println(line)
		re := regexp.MustCompile("[0-9]+")
		allDigits := re.FindAllIndex([]byte(line), -1)
		for _, ab := range allDigits {
			a := ab[0] // index for start of number
			b := ab[1] // index of end of number + 1
			fmt.Printf("Number: %v\n", line[a:b])
			// check horizontals
			if a > 0 {
				fmt.Printf("CheckLeft: %v; ", CheckLeft(a, b, line))
			}
			if b < width {
				fmt.Printf("CheckRight: %v; ", CheckRight(a, b, line))
			}
			// check verticals
			if row > 0 {
				lineAbove := schematic[row-1]
				fmt.Printf("CheckAbove: %v; ", CheckAboveBelow(a, b, lineAbove))
			}
			if row < len(schematic) {
				lineBelow := schematic[row+1]
				fmt.Printf("CheckAbove: %v; ", CheckAboveBelow(a, b, lineBelow))
			}
			fmt.Println()
		}
	}
	fmt.Println(strings.Repeat("=", 140))

}
