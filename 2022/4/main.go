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
	numOverlap := 0
	for _, line := range data {
		elves := strings.Split(line, ",")
		processedElves := map[int][]int{
			0: nil,
			1: nil,
		}

		for index, assignment := range elves {
			elfAssignment := strings.Split(assignment, "-")
			firstElfAssignment, _ := strconv.Atoi(elfAssignment[0])
			secondElfAssignment, _ := strconv.Atoi(elfAssignment[1])
			processedElves[index] = append(processedElves[index], firstElfAssignment)
			processedElves[index] = append(processedElves[index], secondElfAssignment)
		}

		if doAssignmentsOverlap(processedElves) {
			numOverlap = numOverlap + 1
		}
	}
	return numOverlap
}

func doAssignmentsOverlap(assignments map[int][]int) bool {
	result := false

	if assignments[0][0] < assignments[1][0] {
		if assignments[0][1] >= assignments[1][0] {
			result = true
		}
	} else if assignments[0][0] > assignments[1][0] {
		if assignments[1][1] >= assignments[0][0] {
			result = true
		}
	} else {
		result = true
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
