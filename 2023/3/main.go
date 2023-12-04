package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	result := solvePartOne("input.txt")

	log.Printf(`Result: %v`, result)
}

func solvePartOne(filename string) int {
	data := readInput(filename)
	sum := 0

	for lineIndex, line := range data {
		numberRegex := `\d{1,}`
		r, _ := regexp.Compile(numberRegex)
		matches := r.FindAllStringIndex(line, -1)
		for _, match := range matches {
			isValid := isValidPartNumber(match, lineIndex, data)

			if isValid {
				matchedValue := line[match[0]:match[1]]
				partNumber, _ := strconv.Atoi(matchedValue)
				sum = sum + partNumber
			}
		}
	}

	return sum
}

func isValidPartNumber(match []int, lineIndex int, data []string) bool {
	var result []bool
	validPartRegex := `[^\d\.]`
	for i, stringIndex := range match {
		topLine := ""
		bottomLine := ""
		currLine := ""
		initIndex := stringIndex - 1
		endIndex := stringIndex + 2
		if i == 1 {
			// second value in regex match is i + 1, need to drop it back
			initIndex = initIndex - 1
			endIndex = endIndex - 1
		}
		if endIndex > len(data[lineIndex])-1 {
			endIndex = len(data[lineIndex]) - 1
		}
		if stringIndex == 0 {
			initIndex = 0
			currLine = data[lineIndex][initIndex:endIndex]
		} else {
			if stringIndex == len(data[lineIndex])-1 {
				endIndex = len(data[lineIndex]) - 1
			}
			currLine = data[lineIndex][initIndex:endIndex]
		}
		if lineIndex > 0 {
			topLine = data[lineIndex-1][initIndex:endIndex]
		}
		if lineIndex < len(data)-1 {
			bottomLine = data[lineIndex+1][initIndex:endIndex]
		}

		// test each for not(digit or .)
		r, _ := regexp.Compile(validPartRegex)
		topMatches := r.FindAllStringIndex(topLine, -1)
		currMatches := r.FindAllStringIndex(currLine, -1)
		bottomMatches := r.FindAllStringIndex(bottomLine, -1)

		// if any pass check, return true, otherwise false
		result = append(result, len(topMatches) > 0 || len(currMatches) > 0 || len(bottomMatches) > 0)
	}

	return result[0] || result[1]
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
