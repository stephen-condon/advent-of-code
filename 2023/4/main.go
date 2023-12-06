package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := solve("input.txt", false)

	log.Printf(`Result: %v`, result)
}

func solve(filename string, isPartOne bool) int {
	var result int
	data := readInput(filename)
	sum := 0
	cards := initCards(len(data))

	for lineIndex, line := range data {
		_, last, _ := strings.Cut(line, ":")
		winners, ours, _ := strings.Cut(last, " | ")
		winningStrings := strings.Split(trimString(winners), " ")
		ourStrings := strings.Split(trimString(ours), " ")
		winningNumbers := convertStringArrayToInts((winningStrings))
		ourNumbers := convertStringArrayToInts((ourStrings))
		matches := countMatches(winningNumbers, ourNumbers)

		if isPartOne {
			sum = sum + calculateScore(matches)
		} else {
			currentCards := cards[lineIndex]
			for i := lineIndex + 1; i < (matches + lineIndex + 1); i = i + 1 {
				if i < len(cards) {
					cards[i] = cards[i] + currentCards
				} else {
					break
					// if we're trying to add beyond last card, we can kill the loop
				}
			}
		}
	}

	if isPartOne {
		result = sum
	} else {
		result = sumCards(cards)
	}

	return result
}

func sumCards(cards []int) int {
	sum := 0
	for _, numCards := range cards {
		sum = sum + numCards
	}

	return sum
}

func initCards(numCardTypes int) []int {
	cards := []int{}
	for i := 0; i < numCardTypes; i = i + 1 {
		cards = append(cards, 1)
	}

	return cards
}

func calculateScore(count int) int {
	score := 0
	if count > 0 {
		// math.Pow() operates in the float64 space
		score = int(math.Pow(2, float64((count - 1))))
	}
	return score
}

func countMatches(winningNumbers []int, ourNumbers []int) int {
	count := 0
	// for smaller data sets, this is okay - need to think about a better way
	for _, number := range ourNumbers {
		for _, winner := range winningNumbers {
			if number == winner {
				count = count + 1
				break
			}
		}
	}

	return count
}

func convertStringArrayToInts(arr []string) []int {
	var transformed []int
	for _, value := range arr {
		transformedValue, _ := strconv.Atoi(value)
		if transformedValue > 0 {
			// handle extra space for single digit numbers - space converts to zero for integers
			transformed = append(transformed, transformedValue)
		}
	}
	sort.Ints(transformed)
	return transformed
}

func trimString(str string) string {
	return strings.Trim(str, " ")
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
