package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var itemStrings = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var priorityMap = map[string]int{}

func main() {
	result := solveChallenge("input.txt")

	log.Printf(`Result: %v`, result)
}

func solveChallenge(filename string) int {
	sumPriorities := 0
	initializePriorities()
	data := readInput(filename)
	for _, el := range data {
		sumPriorities = sumPriorities + processLine(el)
	}

	return sumPriorities
}

func initializePriorities() {
	items := strings.Split(itemStrings, "")
	for index, el := range items {
		priorityMap[el] = index + 1
	}
}

func processLine(rucksack string) int {
	lineScore := 0
	var compartmentOne = map[string]int{} // not slice, let's do val:index map
	compartmentSize := len(rucksack) / 2
	rucksackItems := strings.Split(rucksack, "")
	isFound := false
	for index, el := range rucksackItems {
		if !isFound {
			if index < compartmentSize {
				compartmentOne[el] = index
			} else {
				// we already have full scope of compartment 1 if we're on the back half of the iteration
				_, found := compartmentOne[el]
				if found {
					isFound = true
					priority := priorityMap[el]
					lineScore = lineScore + priority
				}
			}
		} else {
			break
		}
	}

	return lineScore
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
