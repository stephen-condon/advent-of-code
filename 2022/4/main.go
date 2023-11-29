package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := solveChallenge("input.txt")

	log.Printf(`Result: %v`, result)
}

func solveChallenge(filename string) int {
	data := readInput(filename)
	numFullyContained := 0
	for _, line := range data {
		elves := strings.Split(line, ",")
		processedElves := map[int][]int{
			0: nil,
			1: nil,
		}
		// ["start1-end1", "start2-end2"]
		for index, assignment := range elves {
			elfAssignment := strings.Split(assignment, "-")
			firstElfAssignment, _ := strconv.Atoi(elfAssignment[0])
			secondElfAssignment, _ := strconv.Atoi(elfAssignment[1])
			processedElves[index] = append(processedElves[index], firstElfAssignment)
			processedElves[index] = append(processedElves[index], secondElfAssignment)
			//[start, end]
		}

		if hasFullyContainedPair(processedElves) {
			numFullyContained = numFullyContained + 1
		}
	}
	// split on , to get the two ranges
	// split each range on - to get the two sets of start/end
	// evaluate if one range is fully contained in another
	return numFullyContained
}

func hasFullyContainedPair(assignments map[int][]int) bool {
	result := false
	if assignments[0][0] > assignments[1][0] {
		// second elf is candidate
		if assignments[0][1] <= assignments[1][1] {
			result = true
		}
	} else if assignments[0][0] < assignments[1][0] {
		// first elf is candidate
		if assignments[0][1] >= assignments[1][1] {
			// true
			result = true
		}
	} else {
		result = true
		// start at same point
		// true
	}

	return result
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
		dataSlice = append(dataSlice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataSlice
}
