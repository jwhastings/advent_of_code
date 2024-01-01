// https://adventofcode.com/2023/day/6

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// StrToIntSlice converts a string containing integers delimited by a space into an integer slice
func StrToIntSlice(str string) []int {
	stringSlice := strings.Split(str, " ")
	intSlice := []int{}
	for i := range stringSlice {
		if stringSlice[i] != "" {
			var err error
			integer, err := strconv.Atoi(stringSlice[i])
			if err != nil {
				panic(err)
			}
			intSlice = append(intSlice, integer)
		}
	}
	return intSlice
}

func Distance(timeLimit, timeHeld int) int {
	timeMoving := timeLimit - timeHeld
	boatSpeed := timeHeld
	return timeMoving * boatSpeed
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

func main() {
	filename := "input.txt"
	// filename := "input_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var timeLimit []int
	var distanceRecord []int
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		ints := strings.TrimSpace(input[strings.Index(input, ":")+1:])
		if strings.Index(input, "Time") == 0 {
			timeLimit = StrToIntSlice(ints)
		} else {
			distanceRecord = StrToIntSlice(ints)
		}
	}

	// part one
	waysToWin := make([]int, len(timeLimit))
	for i := range timeLimit {
		for hold := 0; hold <= timeLimit[i]; hold++ {
			distance := Distance(timeLimit[i], hold)
			if distance > distanceRecord[i] {
				waysToWin[i]++
			}
		}
	}

	partOne := 1
	for i := range waysToWin {
		partOne *= waysToWin[i]
	}
	fmt.Printf("Product of ways to win all races (part one): %v\n", partOne)

	// part two
	var timeStr string
	var recordStr string
	for i := range timeLimit {
		timeStr += strconv.Itoa(timeLimit[i])
		recordStr += strconv.Itoa(distanceRecord[i])
	}
	t, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(err)
	}
	r, err := strconv.Atoi(recordStr)
	if err != nil {
		panic(err)
	}

	// upside parabola of form: f(x) = t*x - x^2, where t == time limit (constant), x == time held
	// shift it down by the current record, r:
	// g(x) = f(x) - r = -x^2 + t*x - r
	// set g equal to zero to find values of x where g(x) == 0 <-> f(x) == r. we want the number of
	// integer x values such that g(x) > 0 and they must exist between these two roots of g:
	// -> x = 0.5 * (t +/- sqrt(t^2 - 4*r))
	// this gives us an interval [a, b] such that g(x) >= 0 for all x in [a, b].
	// numbers to the left of a are below the record, numbers to the right of b are below the record,
	// and numbers in between beat the record. we only want integer values of x, so let a' = ceil(a)
	// and let b' = floor(b). The answer is the number of integers between these two numbers (inclusive),
	// i.e. b' - a' + 1
	t2 := IntPow(t, 2)
	discriminant := math.Sqrt(float64(t2 - 4*r))

	a := (float64(t) - discriminant) * 0.5
	a_prime := int(math.Ceil(a))

	b := (float64(t) + discriminant) * 0.5
	b_prime := int(math.Floor(b))

	partTwo := b_prime - a_prime + 1
	fmt.Printf("Number of ways to win big race (part two): %v\n", partTwo)
}
