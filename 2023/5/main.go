package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result1, _ := solve("input.txt", false)

	log.Printf(`Result1: %v`, result1)
}

func solve(filename string, isPartOne bool) (int, int) {
	data := readInput(filename)

	seeds, parsedData := parseInput(data)

	locations := buildAndOrderLocations(seeds, parsedData)

	return locations[0], -1
}

func parseInput(data []string) ([]int, [][][]int) {
	// seeds
	_, seedData, _ := strings.Cut(data[0], ":")
	rawSeeds := strings.Split(strings.Trim(seedData, " "), " ")
	seeds := convertStringArrayToInts(rawSeeds)

	parsedData := [][][]int{}
	colonRegex := `:`
	r, _ := regexp.Compile(colonRegex)
	var rawValues [][]int
	for lineIndex, line := range data {
		colonMatches := r.FindAllStringIndex(line, -1)
		if len(colonMatches) == 0 && lineIndex > 1 {
			// not a header
			if len(line) == 0 || lineIndex == (len(data)-1) {
				// empty line, let's process the section's data, and reset transient data structures
				parsedData = append(parsedData, rawValues)
				rawValues = nil
			} else {
				// process for currentHeader
				rawNumbers := convertStringArrayToInts(strings.Split(line, " "))
				rawValues = append(rawValues, rawNumbers)

			}
		}
	}

	return seeds, parsedData
}

func buildAndOrderLocations(seeds []int, parsedData [][][]int) []int {
	locations := []int{}
	for _, seed := range seeds {
		location := findLocation(seed, parsedData)
		locations = append(locations, location)
	}
	sort.Ints(locations)

	return locations
}

func findLocation(seed int, parsedData [][][]int) int {
	prevValue := seed
	nextValue := -1
	for _, mapSet := range parsedData {
		for _, set := range mapSet {
			if prevValue > set[1] && prevValue <= (set[1]+set[2]) {
				diff := prevValue - set[1]
				nextValue = set[0] + diff
			}
		}
		// get ready for next map; reset nextValue
		if nextValue != -1 {
			prevValue = nextValue
		}
		nextValue = -1
	}

	return prevValue
}

func convertStringArrayToInts(arr []string) []int {
	var transformed []int
	for _, value := range arr {
		transformedValue, _ := strconv.Atoi(value)
		transformed = append(transformed, transformedValue)
	}
	return transformed
}

func readInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var dataSlice []string

	for scanner.Scan() {
		text := scanner.Text()
		dataSlice = append(dataSlice, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataSlice
}
