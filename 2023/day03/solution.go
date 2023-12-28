package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const width int = 140

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var schematic []string
	for scanner.Scan() {
		// get string from input file, line by line
		stringValue := scanner.Text()
		schematic = append(schematic, stringValue)
	}

	fmt.Printf("%v", string(schematic[0][0]))

	file.Close()
}
