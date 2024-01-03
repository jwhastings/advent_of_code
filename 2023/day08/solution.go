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
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		if strings.Index(input, "=") > -1 {
			key := input[:strings.Index(input, " ")]
			values := strings.Split(input[strings.Index(input, "(")+1:strings.Index(input, ")")], ", ")
			leftVal, rightVal := values[0], values[1]
			left[key], right[key] = leftVal, rightVal
		} else if input != "" {
			steps = strings.Split(input, "")
		}
	}

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
}
