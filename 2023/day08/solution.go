// https://adventofcode.com/2023/day/8

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	// filename := "input_test1.txt"
	// filename := "input_test2.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	left := make(map[string]string)
	right := make(map[string]string)
	var steps []string
	var startNodes, endNodes []string
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		if strings.Index(input, "=") > -1 {
			key := input[:strings.Index(input, " ")]
			values := strings.Split(input[strings.Index(input, "(")+1:strings.Index(input, ")")], ", ")
			leftVal, rightVal := values[0], values[1]
			left[key], right[key] = leftVal, rightVal

			if string(key[2]) == "A" {
				startNodes = append(startNodes, key)
			} else if string(key[2]) == "Z" {
				endNodes = append(endNodes, key)
			}
		} else if input != "" {
			steps = strings.Split(input, "")
		}
	}

	// part one
	current := "AAA"
	end := "ZZZ"
	partOne := 0
	for current != end {
		for _, dir := range steps {
			if dir == "L" {
				current = left[current]
			} else if dir == "R" {
				current = right[current]
			}
			partOne++
			if current == end {
				break
			}
		}
	}
	fmt.Printf("Steps taken (part one): %v\n", partOne)

	partTwo := 1
	currentNodes := make([]string, len(startNodes))
	copy(currentNodes, startNodes)
	var nodeCount []int
	for _, node := range currentNodes {
		count := 0
		for string(node[2]) != "Z" {
			for _, dir := range steps {
				if dir == "L" {
					node = left[node]
				} else if dir == "R" {
					node = right[node]
				}
			}
			count++
			if string(node[2]) == "Z" {
				// append number of steps it took to reach the destination
				nodeCount = append(nodeCount, count)
				continue
			}
		}
	}
	// the number of steps to reach the destination for every starting place and
	// the length of step instructions are all prime, so the only way they will
	// all line up and end at their final destination is if they reach a number
	// with all of these values as a divisor. Therefore the answer must be the
	// product of each of these numbers
	for _, mult := range nodeCount {
		partTwo *= mult
	}
	partTwo *= len(steps)
	fmt.Printf("Steps taken (part two): %v\n", partTwo)
}
