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
	result := solvePartOne("input.txt")

	log.Printf(`Result: %v`, result)
}

func solvePartOne(filename string) int {
	data := readInput(filename)
	sum := 0

	for _, line := range data {
		_, last, _ := strings.Cut(line, ":")
		winners, ours, _ := strings.Cut(last, " | ")
		winningStrings := strings.Split(trimString(winners), " ")
		ourStrings := strings.Split(trimString(ours), " ")
		winningNumbers := convertStringArrayToInts((winningStrings))
		ourNumbers := convertStringArrayToInts((ourStrings))
		// log.Println(winningNumbers)
		// log.Println(ourNumbers)
		matches := countMatches(winningNumbers, ourNumbers)
		// log.Println(matches)
		sum = sum + calculateScore(matches)
	}

	return sum
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
