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
	notsymbols string = ".0123456789"
)

type gear struct {
	a int
	b int
}

func IsSymbol(char string) (isSymbol bool, symbol string) {
	if !(strings.ContainsAny(char, notsymbols)) {
		return true, char
	}
	return false, ""
}

func CheckRange(a, b, row int, line string) (adjacent bool, symbol string, loc gear) {
	for i := a; i < b; i++ {
		char := string(line[i])
		symbolCheck, symbol := IsSymbol(char)
		if symbolCheck {
			return true, symbol, gear{i, row}
		}
	}
	return false, "", gear{-1, -1}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var schematic []string
	// var symbols []string
	for scanner.Scan() {
		// get string from input file, line by line, append to slice
		stringValue := scanner.Text()
		schematic = append(schematic, stringValue)
		// iterate over line to identify new symbols
		// for _, char := range stringValue {
		// 	symbols = AppendNewToSlice(symbols, string(char))
		// }
	}
	file.Close()

	var partNumbers []int
	var adjacentSymbol bool
	var gearTable = make(map[gear][]string)
	for row := 0; row < len(schematic); row++ {
		line := schematic[row]
		re := regexp.MustCompile("[0-9]+")
		allDigits := re.FindAllIndex([]byte(line), -1)
		var numStr string
		for _, ab := range allDigits {
			adjacentSymbol = false
			a := ab[0] // index for start of number
			b := ab[1] // index of end of number + 1
			numStr = line[a:b]

			charToLeft := (a > 0)
			charToRight := (b < width)
			charAbove := row > 0
			charBelow := row < len(schematic)-1

			// check horizontals
			if charToLeft {
				adjacent, symbol, location := CheckRange(a-1, a, row, line)
				if adjacent {
					adjacentSymbol = true
					if symbol == "*" {
						gearTable[location] = append(gearTable[location], numStr)
					}
				}
			}
			if charToRight {
				adjacent, symbol, location := CheckRange(b, b+1, row, line)
				if adjacent {
					adjacentSymbol = true
					if symbol == "*" {
						gearTable[location] = append(gearTable[location], numStr)
					}
				}
			}
			// check above
			if charAbove {
				lineAbove := schematic[row-1]

				if charToLeft && charToRight {
					adjacent, symbol, location := CheckRange(a-1, b+1, row-1, lineAbove)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				} else if charToLeft {
					adjacent, symbol, location := CheckRange(a-1, b, row-1, lineAbove)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				} else if charToRight {
					adjacent, symbol, location := CheckRange(a, b+1, row-1, lineAbove)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				} else {
					adjacent, symbol, location := CheckRange(a, b, row-1, lineAbove)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				}
			}
			// check below
			if charBelow {
				lineBelow := schematic[row+1]

				if charToLeft && charToRight {
					adjacent, symbol, location := CheckRange(a-1, b+1, row+1, lineBelow)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				} else if charToLeft {
					adjacent, symbol, location := CheckRange(a-1, b, row+1, lineBelow)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				} else if charToRight {
					adjacent, symbol, location := CheckRange(a, b+1, row+1, lineBelow)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
				} else {
					adjacent, symbol, location := CheckRange(a, b, row+1, lineBelow)
					if adjacent {
						adjacentSymbol = true
						if symbol == "*" {
							gearTable[location] = append(gearTable[location], numStr)
						}
					}
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

	sumPartTwo := 0
	for _, val := range gearTable {
		if len(val) > 1 {
			product := 1
			for _, part := range val {
				partInt, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}
				product *= partInt
			}
			sumPartTwo += product
			// fmt.Printf("%v: %v\n", gear, val)
		}
	}
	fmt.Printf("Sum of gear ratios: %v\n", sumPartTwo)
}
