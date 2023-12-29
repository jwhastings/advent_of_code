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

// ProcessCard increments an array holding the number of times a card has been processed
func ProcessCard(cardNum, matches, totalCards int, timesProcessed []int) []int {
	if matches > 0 {
		lower := cardNum + 1
		upper := min(totalCards, lower+matches)
		for i := lower; i < upper; i++ {
			timesProcessed[i]++
		}
	}
	return timesProcessed
}

// IntPow returns n to the m-th power for only integer inputs
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

// NumMatches returns the number of matches found between winning numbers ones we have
func NumMatches(winners, have []string) (matches int) {
	for _, winner := range winners {
		if slices.Contains(have, winner) {
			// fmt.Printf("%v;", winner)
			matches++
		}
	}
	return matches
}

// CalcPoints returns the number of points received for a given number of matches
func CalcPoints(matches int) (points int) {
	if matches > 0 {
		points = IntPow(2, matches-1)
	}
	return points
}

func main() {
	filename := "input.txt"
	// filename := "input_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var cards []string
	for scanner.Scan() {
		// get string from input file, line by line
		card := scanner.Text()
		cards = append(cards, card)
	}
	file.Close()

	var sumPartOne, sumPartTwo int
	var timesProcessed = make([]int, len(cards))
	for count := range timesProcessed {
		timesProcessed[count] = 1 // default all card counts to 1
	}

	for cardNum, card := range cards {
		// split between winnings numbers and numbers we have
		cardSplit := strings.Split(card[strings.Index(card, ":")+1:], "|")
		// trim leading/trailing space, replace double spaces with a single space, split into slice delimited by space
		winningNums := strings.Split(strings.ReplaceAll(strings.TrimSpace(cardSplit[0]), "  ", " "), " ")
		slices.Sort(winningNums)
		haveNums := strings.Split(strings.ReplaceAll(strings.TrimSpace(cardSplit[1]), "  ", " "), " ")
		slices.Sort(haveNums)

		// part one
		matches := NumMatches(winningNums, haveNums)
		points := CalcPoints(matches)
		sumPartOne += points

		// part two
		repeat := timesProcessed[cardNum]
		// for every copy of a card
		for i := 1; i <= repeat; i++ {
			timesProcessed = ProcessCard(cardNum, matches, len(cards), timesProcessed)
			sumPartTwo++ // running total of how many cards have been counted
		}
	}

	fmt.Printf("Total points: %v\n", sumPartOne)
	fmt.Printf("Total cards processed: %v\n", sumPartTwo)
}
