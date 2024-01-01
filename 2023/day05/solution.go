// https://adventofcode.com/2023/day/5

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

type destination struct {
	group   string
	data    []string
	mapping map[int]int
}

// StrToIntSlice converts a string containing integers delimited by a space into an integer slice
func StrToIntSlice(str string) []int {
	str = strings.TrimSpace(strings.ReplaceAll(str, "  ", " "))
	stringSlice := strings.Split(str, " ")
	intSlice := make([]int, len(stringSlice))
	for i := range stringSlice {
		var err error
		intSlice[i], err = strconv.Atoi(stringSlice[i])
		if err != nil {
			panic(err)
		}
	}
	return intSlice
}

// AddDataToMap adds key-value pairs from its data if the key falls within the specified data
// Update key to be a range of values (slice) instead of a single int
func AddDataToMap(keyRange []int, d destination) destination {
	d.mapping = make(map[int]int)
	for _, input := range d.data {
		intSlice := StrToIntSlice(input)
		destinationStart := intSlice[0]
		sourceStart := intSlice[1]
		rangeLength := intSlice[2]
		sourceRange := []int{sourceStart, sourceStart + rangeLength - 1}
		// given key range falls within data range
		// if key >= sourceStart && key <= sourceStart+rangeLength {
		if RangeOverlap(keyRange, sourceRange) {
			intersection := RangeIntersect(keyRange, sourceRange)
			keysOverlapping := RangeSequence(intersection)
			mapping := []int{keyRange[0] - sourceRange[0] + destinationStart, keyRange[1] - source}
			// for _, key := range keysOverlapping {
			// 	d.mapping[key] = key - sourceStart + destinationStart
			// }
			// return d // mapping exists in data
		}
	}
	return d // mapping does not exist in data
}

// RangeOverlap takes two integer slices with two elements containing the start and end
// of a range (both inclusive) and returns whether they overlap
func RangeOverlap(x, y []int) bool {
	return x[0] <= y[1] && y[0] <= x[1]
}

// RangeIntersect takes two integer slices with two elements containing the start and end
// of a range (both inclusive) and returns a new range of their intersection
func RangeIntersect(x, y []int) []int {
	if RangeOverlap(x, y) {
		return []int{
			max(x[0], y[0]),
			min(x[1], y[1]),
		}
	}
	return []int{}
}

// RangeSequence takes an integer slice with two elements containing the start and end of
// a range (both inclusive) and returns a new slice containing all integers within the range
func RangeSequence(x []int) (seq []int) {
	seq = make([]int, x[1]-x[0]+1)
	i := 0
	for j := x[0]; j < x[1]+1; j++ {
		seq[i] = j
		i++
	}
	return seq
}

// RetrieveDestination returns the value for a given key and map with a destination struct.
// If the key doesn't exist in the map, the key itself is returned
func RetrieveDestination(keyRange []int, d destination) int {
	// value, exists := d.mapping[key]
	// if exists {
	// 	return value
	// }
	// return key
	return 1 // temporary for debugging
}

// AddAndRetrieve
func AddAndRetrieve(keyRange []int, d destination) int {
	d = AddDataToMap(keyRange, d)
	value := RetrieveDestination(keyRange, d)
	return value
}

func main() {
	// filename := "input.txt"
	filename := "input_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var seedString, group string
	seedToSoil := destination{}
	soilToFertilizer := destination{}
	fertilizerToWater := destination{}
	waterToLight := destination{}
	lightToTemperature := destination{}
	temperatureToHumidity := destination{}
	humidityToLocation := destination{}

	// input processing
	reGroup := regexp.MustCompile(`\w+\-to\-\w+ map\:`)
	reMap := regexp.MustCompile("^[0-9]+ [0-9]+ [0-9]+( )*$")
	for scanner.Scan() {
		// get string from input file, line by line
		input := scanner.Text()
		var groupMatch, mapMatch bool
		if input != "" {
			mapMatch = reMap.MatchString(input)
			if !mapMatch {
				groupMatch = reGroup.MatchString(input)
			}
		}

		if mapMatch {
			// add source/destination data to maps depending on group
			switch group {
			case "seed-to-soil":
				seedToSoil.group = group
				seedToSoil.data = append(seedToSoil.data, input)
			case "soil-to-fertilizer":
				soilToFertilizer.group = group
				soilToFertilizer.data = append(soilToFertilizer.data, input)
			case "fertilizer-to-water":
				fertilizerToWater.group = group
				fertilizerToWater.data = append(fertilizerToWater.data, input)
			case "water-to-light":
				waterToLight.group = group
				waterToLight.data = append(waterToLight.data, input)
			case "light-to-temperature":
				lightToTemperature.group = group
				lightToTemperature.data = append(lightToTemperature.data, input)
			case "temperature-to-humidity":
				temperatureToHumidity.group = group
				temperatureToHumidity.data = append(temperatureToHumidity.data, input)
			case "humidity-to-location":
				humidityToLocation.group = group
				humidityToLocation.data = append(humidityToLocation.data, input)
			}
		} else if groupMatch {
			group = strings.TrimSpace(input[:strings.Index(input, " ")])
		} else if strings.Contains(input, "seeds: ") {
			seedString = input[strings.Index(input, ":")+1:]
		}
	}

	seedInit := StrToIntSlice(seedString)
	// seedLocation := make([]int, len(seedInit))
	// for i, seed := range seedInit {
	// 	soil := AddAndRetrieve([]int{seed, seed}, seedToSoil)
	// 	fertilizer := AddAndRetrieve([]int{soil, soil}, soilToFertilizer)
	// 	water := AddAndRetrieve([]int{fertilizer, fertilizer}, fertilizerToWater)
	// 	light := AddAndRetrieve([]int{water, water}, waterToLight)
	// 	temperature := AddAndRetrieve([]int{light, light}, lightToTemperature)
	// 	humidity := AddAndRetrieve([]int{temperature, temperature}, temperatureToHumidity)
	// 	location := AddAndRetrieve([]int{humidity, humidity}, humidityToLocation)
	// 	seedLocation[i] = location
	// }

	// var partOne int
	// for i := range seedLocation {
	// 	if i == 1 {
	// 		partOne = seedLocation[i] // default nearest location to first location
	// 	} else if seedLocation[i] < partOne {
	// 		partOne = seedLocation[i] // assign new nearest location
	// 	}
	// }
	// fmt.Printf("Nearest location (part 1): %v\n", partOne)

	var seedRanges [][]int
	for i := 0; i < len(seedInit)/2+1; i += 2 {
		start := seedInit[i]
		length := seedInit[i+1]
		seedRanges = append(seedRanges, []int{start, start + length - 1})
	}

	// newSeedLocation := make([]int, len(seedRanges))
	for _, seed := range seedRanges {
		soil := AddAndRetrieve(seed, seedToSoil)
		fmt.Printf("%v", soil)
		// fertilizer := AddAndRetrieve(soil, soilToFertilizer)
		// water := AddAndRetrieve(fertilizer, fertilizerToWater)
		// light := AddAndRetrieve(water, waterToLight)
		// temperature := AddAndRetrieve(light, lightToTemperature)
		// humidity := AddAndRetrieve(temperature, temperatureToHumidity)
		// location := AddAndRetrieve(humidity, humidityToLocation)
		// newSeedLocation[i] = location
	}
	// var partTwo int
	// for i := range newSeedLocation {
	// 	if i == 1 {
	// 		partTwo = newSeedLocation[i]
	// 	} else if newSeedLocation[i] < partTwo {
	// 		partTwo = newSeedLocation[i]
	// 	}
	// }
	// fmt.Printf("Nearest location (part 2): %v\n", partTwo)
}
