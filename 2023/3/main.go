package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	result := solvePartTwo("input.txt")

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

func solvePartTwo(filename string) int {
	data := readInput(filename)
	sumGearRatios := 0
	linesNumberRegexMatches := loadNumberMatches(data)

	for lineIndex, line := range data {
		matches := gearMatches(line)
		for _, match := range matches {
			gearIndex := match[0]
			topMatches := [][]int{}
			bottomMatches := [][]int{}
			topData := ""
			bottomData := ""
			if lineIndex != 0 {
				topMatches = linesNumberRegexMatches[lineIndex-1]
				topData = data[lineIndex-1]
			}
			if lineIndex != len(data)-1 {
				bottomMatches = linesNumberRegexMatches[lineIndex+1]
				bottomData = data[lineIndex+1]
			}
			surroundingMatches := [][][]int{
				topMatches,
				linesNumberRegexMatches[lineIndex],
				bottomMatches,
			}
			surroundingData := []string{
				topData,
				data[lineIndex],
				bottomData,
			}
			surroundingNumbers := calculateSurroundingNumbers(gearIndex, surroundingMatches, surroundingData)
			if len(surroundingNumbers) == 2 {
				// make gear ratio
				sumGearRatios = sumGearRatios + (surroundingNumbers[0] * surroundingNumbers[1])
			}
		}
	}

	return sumGearRatios

}

func calculateSurroundingNumbers(gearIndex int, surroundingMatches [][][]int, surroundingData []string) []int {
	var surroundingNumbers []int
	for index, line := range surroundingData {
		for _, match := range surroundingMatches[index] {
			// for each line, if 1st match index
			if match[0] >= (gearIndex-1) && match[0] <= (gearIndex+1) {
				// matched
				number, _ := strconv.Atoi(line[match[0]:match[1]])
				surroundingNumbers = append(surroundingNumbers, number)
			} else if (match[1]-1) >= (gearIndex-1) && (match[1]-1) <= (gearIndex+1) {
				// matched
				number, _ := strconv.Atoi(line[match[0]:match[1]])
				surroundingNumbers = append(surroundingNumbers, number)
			}
		}
	}

	return surroundingNumbers
}

func loadNumberMatches(data []string) [][][]int {
	var linesNumberRegexMatches [][][]int

	for _, line := range data {
		numberRegex := `\d{1,}`
		numberR, _ := regexp.Compile(numberRegex)
		matches := numberR.FindAllStringIndex(line, -1)
		linesNumberRegexMatches = append(linesNumberRegexMatches, matches)
	}
	return linesNumberRegexMatches
}

func gearMatches(line string) [][]int {
	gearRegex := `\*`
	gearR, _ := regexp.Compile(gearRegex)
	matches := gearR.FindAllStringIndex(line, -1)
	return matches
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
