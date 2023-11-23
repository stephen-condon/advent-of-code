package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var outcomeMap = map[string]int{
	"win":  6,
	"draw": 3,
	"loss": 0,
}
var choiceMap = map[string]int{
	"r": 1,
	"p": 2,
	"s": 3,
}

var victoryMap = map[string]string{
	"r": "s",
	"p": "r",
	"s": "p",
}

var symbolTranslation = map[string]string{
	"A": "r",
	"B": "p",
	"C": "s",
}

var resultTranslation = map[string]string{
	"X": "loss",
	"Y": "draw",
	"Z": "win",
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
		totalScore = totalScore + processRound(value)
	}

	return totalScore
}

func processRound(round string) int {
	roundSlice := strings.Split(round, " ")
	rawTheirs := roundSlice[0]
	rawResult := roundSlice[1]
	theirs := translateSymbol(rawTheirs)
	outcome := translateResult(rawResult)

	var ours string

	if outcome == "draw" {
		ours = theirs
	} else if outcome == "win" {
		// reverse lookup in victoryMap
		ours = victoryReverseLookup(theirs)
	} else {
		ours = victoryMap[theirs]
	}

	score := outcomeMap[outcome] + choiceMap[ours]

	return score
}

func victoryReverseLookup(loser string) string {
	var symbol string
	for key, value := range victoryMap {
		if value == loser {
			symbol = key
		}
	}
	return symbol
}

func translateSymbol(symbol string) string {
	return symbolTranslation[symbol]
}

func translateResult(result string) string {
	return resultTranslation[result]
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
