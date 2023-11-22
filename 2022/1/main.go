package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var filename = "input.txt"

func main() {
	data := readInput()
	maxElfCaloriesIndex, maxElfCalories, topThreeSum := processData(data)
	log.Printf(`Index %v: %v Calories`, maxElfCaloriesIndex, maxElfCalories)
	log.Printf(`Top 3 Elves: %v`, topThreeSum)
}

func processData(data []string) (int, int, int) {
	var elfCalories []int

	currentElfCalories := 0
	maxElfCalories := 0
	maxElfCaloriesIndex := -1

	for _, value := range data {
		if len(value) == 0 {
			elfCalories = append(elfCalories, currentElfCalories)
			if currentElfCalories > maxElfCalories {
				maxElfCalories = currentElfCalories
				maxElfCaloriesIndex = len(elfCalories)
			}
			currentElfCalories = 0
		} else {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			currentElfCalories = currentElfCalories + intValue
		}
	}

	sort.Ints(elfCalories)

	topThreeSum := elfCalories[len(elfCalories)-1] + elfCalories[len(elfCalories)-2] + elfCalories[len(elfCalories)-3]

	return maxElfCaloriesIndex, maxElfCalories, topThreeSum
}

func readInput() []string {
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
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataSlice
}
