package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var numberTextLookup map[string]int
var regexString []string

func main() {
	result := solvePartTwo("input.txt")

	log.Printf(`Result: %v`, result)
}

func solvePartOne(filename string) int {
	regexString = []string{`\d`}

	data := readInput(filename)

	return calculateSumCalibrationValues(data)
}

func solvePartTwo(filename string) int {
	initNumberTextLookup()
	regexString = []string{`\d`, `(?:one)`, `(?:two)`, `(?:three)`, `(?:four)`, `(?:five)`, `(?:six)`, `(?:seven)`, `(?:eight)`, `(?:nine)`}

	data := readInput(filename)

	return calculateSumCalibrationValues(data)
}

func calculateSumCalibrationValues(data []string) int {
	sum := 0

	for _, line := range data {
		matches := [][]int{}
		// need to iterate through regex, and dynamically build matches, based on example in bug.txt
		// there may be a better regex way to handle this, but I haven't figured it out
		for _, value := range regexString {
			r, _ := regexp.Compile(value)
			individualMatches := r.FindAllStringIndex(line, -1)
			matches = append(matches, individualMatches...)
		}
		sort.SliceStable(matches, func(i, j int) bool { return matches[i][0] < matches[j][0] })

		firstMatch := line[matches[0][0]:matches[0][1]]
		lastMatch := line[matches[len(matches)-1][0]:matches[len(matches)-1][1]]
		var firstDigit, secondDigit int

		if len(firstMatch) > 1 {
			// process string number
			firstDigit = translateNumberText(firstMatch)
		} else {
			firstDigit, _ = strconv.Atoi(line[matches[0][0]:matches[0][1]])
		}
		if len(lastMatch) > 1 {
			// process string number
			secondDigit = translateNumberText(lastMatch)
		} else {
			secondDigit, _ = strconv.Atoi(line[matches[len(matches)-1][0]:matches[len(matches)-1][1]])
		}

		calibrationValue := (firstDigit * 10) + secondDigit

		sum = sum + calibrationValue
	}

	return sum
}

func initNumberTextLookup() {
	numberTextLookup = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
}

func translateNumberText(text string) int {
	return numberTextLookup[text]
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
