package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var gameMaximums GameValues
var gamePossible bool

func main() {
	result := solvePartTwo("input.txt")

	log.Printf(`Result: %v`, result)
}

func solvePartOne(filename string) int {
	total := 0
	initGameMaximums()
	data := readInput(filename)

	for _, value := range data {
		total = total + processGame(value)
	}

	return total
}

func solvePartTwo(filename string) int {
	sum := 0
	data := readInput(filename)

	for _, game := range data {
		gameMinimums := GameValues{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		initialBreakdown := strings.Split(game, ": ")

		samples := strings.Split(initialBreakdown[1], "; ")

		for _, sample := range samples {
			colorData := strings.Split(sample, ", ")
			for _, value := range colorData {
				splitValue := strings.Split(value, " ")
				prevIndexValue, _ := strconv.Atoi(splitValue[0])
				if gameMinimums[splitValue[1]] < prevIndexValue {
					gameMinimums[splitValue[1]] = prevIndexValue
				}
			}
		}

		sum = sum + (gameMinimums["blue"] * gameMinimums["red"] * gameMinimums["green"])
	}

	return sum
}

func processGame(game string) int {
	gameValue := 0
	gamePossible = true
	initialBreakdown := strings.Split(game, ": ")

	samples := strings.Split(initialBreakdown[1], "; ")

	parseSamples(samples)

	if gamePossible {
		gameValue = getGameId(initialBreakdown[0])
	}

	return gameValue
}

func parseSamples(samples []string) {
	for _, sample := range samples {
		if !gamePossible {
			break
		}
		colorData := strings.Split(sample, ", ")
		for _, value := range colorData {
			splitValue := strings.Split(value, " ")
			prevIndexValue, _ := strconv.Atoi(splitValue[0])
			if gameMaximums[splitValue[1]] < prevIndexValue {
				gamePossible = false
				break
			}
		}
	}
}

func getGameId(gameMetadata string) int {
	idBreakdown := strings.Split(gameMetadata, " ")
	gameId, _ := strconv.Atoi(idBreakdown[1])

	return gameId
}

func initGameMaximums() {
	gameMaximums = GameValues{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
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

type GameValues map[string]int
