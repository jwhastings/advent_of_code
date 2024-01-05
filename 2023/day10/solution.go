// https://adventofcode.com/2023/day/10

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	distance int
	tile     string
	location [2]int
	top      *node
	bottom   *node
	left     *node
	right    *node
}

func PrintBoard(d [][]node) {
	fmt.Printf("\n")
	for _, row := range d {
		for _, node := range row {
			fmt.Printf(node.tile)
		}
		fmt.Printf("\n")
	}
}

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

	stringMap := make(map[string]string)
	old := []string{"|", "-", "L", "J", "7", "F", ".", "S"}
	new := []string{"║", "═", "╚", "╝", "╗", "╔", ".", "S"}
	for i := range old {
		stringMap[old[i]] = new[i]
	}

	var diagram [][]node
	var startPos []int
	var x, y int
	i := 0
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		split := strings.Split(input, "")
		nodeRow := make([]node, len(split))
		for j, s := range split {
			newNode := &node{
				distance: 0,
				tile:     stringMap[s],
				location: [2]int{i, j},
			}
			nodeRow[j] = *newNode
		}

		diagram = append(diagram, nodeRow)
		if strings.Contains(input, "S") {
			startPos = []int{len(diagram) - 1, strings.Index(input, "S")}
			fmt.Printf("Starting position: %v\n", startPos)
		}
		i++
	}
	x = len(diagram)
	y = len(diagram[0])
	PrintBoard(diagram)
	fmt.Printf("\n%v * %v board\n", x, y)
}
