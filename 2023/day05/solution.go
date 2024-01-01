// https://adventofcode.com/2023/day/5

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// data structure to hold intervals
type interval struct {
	interval []int
}

// CreateInterval creates an interval from a start (a) and end (b) position
func CreateInterval(a, b int) interval {
	return interval{[]int{a, b}}
}

// a accesses the lower end of an interval
func (i interval) a() int {
	return i.interval[0]
}

// b accesses the upper end of an interval
func (i interval) b() int {
	return i.interval[1]
}

// data structure to hold mapping information from one source to the next
type mapping struct {
	group       string
	data        []string
	source      []interval
	destination []interval
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

// GetDestination returns destination intervals for the next mapping
func GetDestination(keyIntervals []interval, m mapping) []interval {
	// while we have key intervals to check against the mapping data
	for len(keyIntervals) > 0 {
		found := false         // initialize a matching key mapping to false
		key := keyIntervals[0] // grab our first key interval
		for _, input := range m.data {
			intSlice := StrToIntSlice(input)
			mappingStart := intSlice[0] // start of the mapping interval
			sourceStart := intSlice[1]  // start of the source interval
			diff := sourceStart - mappingStart
			rangeLength := intSlice[2]
			source := CreateInterval(sourceStart, sourceStart+rangeLength-1)
			// given key falls within data range
			if IntervalOverlap(key, source) {
				intersection := IntervalIntersect(key, source)
				// intersection is a subset of the key
				if key.a() != intersection.a() || key.b() != intersection.b() {
					// lower end of key is below start of intersection
					if key.a() < intersection.a() {
						keyLower := CreateInterval(key.a(), intersection.a()-1)
						keyIntervals = append(keyIntervals, keyLower)
					}
					// upper end of key is above end of intersection
					if key.b() > intersection.b() {
						keyUpper := CreateInterval(intersection.b()+1, key.b())
						keyIntervals = append(keyIntervals, keyUpper)
					}
				}
				mapping := CreateInterval(intersection.a()-diff, intersection.b()-diff)
				m.source = append(m.source, intersection)
				m.destination = append(m.destination, mapping)
				// key mapping found, remove from the ones we need to check
				found = true
				keyIntervals = slices.Delete(keyIntervals, 0, 1)
				break
			}
		}
		// iterated through all mapping data without a match
		if !found {
			m.source = append(m.source, key)
			m.destination = append(m.destination, key)
			keyIntervals = slices.Delete(keyIntervals, 0, 1)
		}
	}
	return m.destination
}

// IntervalOverlap takes two integer slices with two elements containing the start and end
// of a range (both inclusive) and returns whether they overlap
func IntervalOverlap(x, y interval) bool {
	return x.a() <= y.b() && y.a() <= x.b()
}

// IntervalIntersect takes two integer slices with two elements containing the start and end
// of a range (both inclusive) and returns a new range of their intersection
func IntervalIntersect(x, y interval) interval {
	if IntervalOverlap(x, y) {
		return interval{[]int{
			max(x.a(), y.a()),
			min(x.b(), y.b())},
		}
	}
	return interval{}
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

	var seedString, group string
	seedToSoil := mapping{}
	soilToFertilizer := mapping{}
	fertilizerToWater := mapping{}
	waterToLight := mapping{}
	lightToTemperature := mapping{}
	temperatureToHumidity := mapping{}
	humidityToLocation := mapping{}

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
	seeds := make([]interval, len(seedInit))
	for i, seed := range seedInit {
		seeds[i] = CreateInterval(seed, seed)
	}
	soil := GetDestination(seeds, seedToSoil)
	fertilizer := GetDestination(soil, soilToFertilizer)
	water := GetDestination(fertilizer, fertilizerToWater)
	light := GetDestination(water, waterToLight)
	temperature := GetDestination(light, lightToTemperature)
	humidity := GetDestination(temperature, temperatureToHumidity)
	locations := GetDestination(humidity, humidityToLocation)

	var partOne int
	for i, location := range locations {
		if i == 0 {
			partOne = location.a() // default nearest location to first location
		} else if location.a() < partOne {
			partOne = location.a() // assign new nearest location
		}
	}
	fmt.Printf("Nearest location (part 1): %v\n", partOne)

	var seedspt2 = make([]interval, len(seedInit)/2)
	j := 0
	for i := 0; i < len(seedInit); i += 2 {
		start := seedInit[i]
		length := seedInit[i+1]
		seedspt2[j] = CreateInterval(start, start+length-1)
		j++
	}

	soil = GetDestination(seedspt2, seedToSoil)
	fertilizer = GetDestination(soil, soilToFertilizer)
	water = GetDestination(fertilizer, fertilizerToWater)
	light = GetDestination(water, waterToLight)
	temperature = GetDestination(light, lightToTemperature)
	humidity = GetDestination(temperature, temperatureToHumidity)
	locations = GetDestination(humidity, humidityToLocation)

	var partTwo int
	for i, location := range locations {
		if i == 0 {
			partTwo = location.a() // default nearest location to first location
		} else if location.a() < partTwo {
			partTwo = location.a() // assign new nearest location
		}
	}
	fmt.Printf("Nearest location (part 2): %v\n", partTwo)
}
