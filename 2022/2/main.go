package main

import (
	"bufio"
	"log"
	"os"
)

var shortcutMap = map[string]int{
	"A X": 4, // rock-rock draw
	"A Y": 8, // rock-paper win 8
	"A Z": 3, // rock-scissors loss
	"B X": 1, // paper-rock loss 1
	"B Y": 5, // paper-paper draw
	"B Z": 9, // paper-scissors win
	"C X": 7, // scissors-rock win
	"C Y": 2, // scissors-paper loss
	"C Z": 6, // scissors-scissors draw
}

func main() {

	total := solveChallenge("input.txt")

	log.Printf(`Total Score: %v`, total)
}

func solveChallenge(filename string) int {
	data := readInput(filename)
	return processData(data)
}

func processData(data []string) int {

	var totalScore = 0

	for _, value := range data {
		totalScore = totalScore + shortcutMap[value]
	}

	return totalScore
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
