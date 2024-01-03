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
	cards     []string
	bid       int
	typename  string
	typerank  int
	typenameJ string
	typerankJ int
}

func CreateHand(cards []string, bid int) hand {
	h := hand{cards, bid, "", -1, "", -1}
	// part one rules
	counts := CountCards(h.cards)
	h.typename, h.typerank = RankCards(counts)

	// part two rules with joker wild card
	jokerCount := 0
	var cardsNoJoker []string
	for _, card := range h.cards {
		if card == "J" {
			jokerCount++
		} else {
			cardsNoJoker = append(cardsNoJoker, card)
		}
	}

	if jokerCount == 0 {
		// if no jokers, hand type and rank is the same as in part one
		h.typenameJ, h.typerankJ = h.typename, h.typerank
	} else if jokerCount == 5 {
		// hand is all jokers, can't substitute in for other cards
		h.typenameJ, h.typerankJ = "five of a kind", 6
	} else {
		h.typenameJ, h.typerankJ = BestJokerHand(jokerCount, cardsNoJoker)
	}
	return h
}

func CountCards(cards []string) []int {
	var counts []int
	for _, count := range CardMap(cards) {
		counts = append(counts, count)
	}
	return counts
}

func CardMap(cards []string) map[string]int {
	cardMap := make(map[string]int)
	for _, card := range cards {
		cardMap[card]++
	}
	return cardMap
}

func RankCards(counts []int) (string, int) {
	// reverse sort so most frequent card is first
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	if counts[0] == 5 {
		return "five of a kind", 6
	} else if counts[0] == 4 {
		return "four of a kind", 5
	} else if counts[0] == 3 && counts[1] == 2 {
		return "full house", 4
	} else if counts[0] == 3 {
		return "three of a kind", 3
	} else if counts[0] == 2 && counts[1] == 2 {
		return "two pair", 2
	} else if counts[0] == 2 {
		return "one pair", 1
	}
	return "high card", 0
}

func BestJokerHand(jokerCount int, cardsNoJoker []string) (string, int) {
	bestType, bestRank := "", 0
	cardMap := CardMap(cardsNoJoker)
	// for each unique non-joker card in the hand
	for card := range cardMap {
		// can only replace as many cards as we have jokers in the hand
		replaceCount := jokerCount
		// copy our joker-less hand into a temporary string slice
		tempCards := make([]string, len(cardsNoJoker))
		copy(tempCards, cardsNoJoker)
		for ; replaceCount > 0; replaceCount-- {
			tempCards = append(tempCards, card)
		}
		jokerType, jokerRank := RankCards(CountCards(tempCards))
		if jokerRank > bestRank {
			bestRank = jokerRank
			bestType = jokerType
		}
	}
	return bestType, bestRank
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

	var cardRank map[string]int
	cardRank = make(map[string]int)
	cardLabels := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	for i := range cardLabels {
		cardRank[cardLabels[i]] = i
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
	}

	// part one
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

	var partOne int
	for i, hand := range allHands {
		partOne += hand.bid * (i + 1)
	}
	fmt.Printf("Total winnings (part one): %v\n", partOne)

	// part two
	var cardRankJ map[string]int
	cardRankJ = make(map[string]int)
	// move joker to worst card
	cardLabelsJ := []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	for i := range cardLabelsJ {
		cardRankJ[cardLabelsJ[i]] = i
	}
	sort.SliceStable(allHands, func(i, j int) bool {
		if allHands[i].typerankJ == allHands[j].typerankJ {
			for k := range allHands[i].cards {
				cardI := allHands[i].cards[k]
				cardJ := allHands[j].cards[k]
				if cardI == cardJ {
					continue
				}
				return cardRankJ[cardI] < cardRankJ[cardJ]
			}
		}
		return allHands[i].typerankJ < allHands[j].typerankJ
	})

	var partTwo int
	for i, hand := range allHands {
		partTwo += hand.bid * (i + 1)
	}
	fmt.Printf("Total winnings with jokers (part two): %v\n", partTwo)
}
