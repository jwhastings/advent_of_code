// https://adventofcode.com/2023/day/4

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var sumPartOne int
	for scanner.Scan() {
		// get string from input file, line by line
		card := scanner.Text()
		cardSplit := strings.Split(card[strings.Index(card, ":")+1:], "|")
		// trim leading/trailing space, replace double spaces with a single space, split into slice delimited by space
		winningNums := strings.Split(strings.ReplaceAll(strings.TrimSpace(cardSplit[0]), "  ", " "), " ")
		slices.Sort(winningNums)
		haveNums := strings.Split(strings.ReplaceAll(strings.TrimSpace(cardSplit[1]), "  ", " "), " ")
		slices.Sort(haveNums)

		matches := 0
		points := 0
		// fmt.Printf("Winning Numbers: %v; Have Numbers: %v\n", winningNums, haveNums)
		for _, winner := range winningNums {
			if slices.Contains(haveNums, winner) {
				// fmt.Printf("%v;", winner)
				matches++
			}
		}
		if matches > 0 {
			points = IntPow(2, matches-1)
		}
		// fmt.Printf("\nMatches: %v; Points: %v\n", matches, points)
		sumPartOne += points
	}

	file.Close()

	fmt.Printf("Total points: %v\n", sumPartOne)
}
