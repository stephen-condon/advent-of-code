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
	var elfGroup []string
	for _, el := range data {
		elfGroup = append(elfGroup, el)
		if len(elfGroup) == 3 {
			badgePriority := processGroup(elfGroup)
			log.Printf(`Badge Priority: %v`, badgePriority)
			sumPriorities = sumPriorities + badgePriority
			//reset elfGroup
			elfGroup = nil
		}

	}

	return sumPriorities
}

func initializePriorities() {
	items := strings.Split(itemStrings, "")
	for index, el := range items {
		priorityMap[el] = index + 1
	}
}

func processGroup(elfGroup []string) int {
	counts := map[string]int{}
	badgePriority := -1
	for _, el := range elfGroup {
		newCounts, badge := processElf(el, counts)
		counts = newCounts
		log.Printf(`Result: %v`, counts)
		log.Printf(`Badge: %v`, badge)
		log.Printf(`Priority: %v`, priorityMap[badge])
		if badge != "" {
			badgePriority = priorityMap[badge]
			break
		}
	}

	return badgePriority
}

func processElf(rucksack string, counts map[string]int) (map[string]int, string) {
	itemsProcessed := map[string]int{}
	badge := ""
	rucksackItems := strings.Split(rucksack, "")
	for _, el := range rucksackItems {
		_, success := itemsProcessed[el]
		if !success {
			itemsProcessed[el] = 1
			counts[el] = counts[el] + 1
			if counts[el] == 3 {
				badge = el
				break
			}
		}
	}

	return counts, badge
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
