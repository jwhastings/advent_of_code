// https://adventofcode.com/2023/day/7

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards    []string
	bid      int
	typename string
	typerank int
}

func CreateHand(cards []string, bid int) hand {
	h := hand{cards, bid, "", -1}

	typemap := make(map[string]int)
	for _, card := range h.cards {
		typemap[card]++
	}
	var counts []int
	for _, count := range typemap {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	if counts[0] == 5 {
		h.typename = "five of a kind"
		h.typerank = 6
	} else if counts[0] == 4 {
		h.typename = "four of a kind"
		h.typerank = 5
	} else if counts[0] == 3 && counts[1] == 2 {
		h.typename = "full house"
		h.typerank = 4
	} else if counts[0] == 3 {
		h.typename = "three of a kind"
		h.typerank = 3
	} else if counts[0] == 2 && counts[1] == 2 {
		h.typename = "two pair"
		h.typerank = 2
	} else if counts[0] == 2 {
		h.typename = "one pair"
		h.typerank = 1
	} else {
		h.typename = "high card"
		h.typerank = 0
	}

	return h
}

var cardRank map[string]int

func main() {
	filename := "input.txt"
	// filename := "input_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cardRank = make(map[string]int)
	cardlabels := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	for i := range cardlabels {
		cardRank[cardlabels[i]] = i
	}

	allHands := []hand{}
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		cards := strings.Split(strings.TrimSpace(input[:strings.Index(input, " ")]), "")
		bid, err := strconv.Atoi(strings.TrimSpace(input[strings.Index(input, " ")+1:]))
		if err != nil {
			panic(err)
		}

		hand := CreateHand(cards, bid)
		allHands = append(allHands, hand)
		sort.SliceStable(allHands, func(i, j int) bool {
			if allHands[i].typerank == allHands[j].typerank {
				for k := range allHands[i].cards {
					cardI := allHands[i].cards[k]
					cardJ := allHands[j].cards[k]
					if cardI == cardJ {
						continue
					}
					return cardRank[cardI] < cardRank[cardJ]
				}
			}
			return allHands[i].typerank < allHands[j].typerank
		})

	}

	var partOne int
	for i, hand := range allHands {
		partOne += hand.bid * (i + 1)
	}
	fmt.Printf("Total winnings (part one): %v", partOne)
}
