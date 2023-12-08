package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result1, result2 := solve("input.txt", false)

	log.Printf(`Result1: %v; Result2: %v`, result1, result2)
}

func solve(filename string, isPartOne bool) (int, int) {
	data := readInput(filename)

	time, distance, compressedTime, compressedDistance := parseInput(data)

	part1 := evaluateRaces(time, distance)
	part2 := evaluateRace(compressedTime, compressedDistance)

	return part1, part2
}

func evaluateRaces(time []int, distance []int) int {
	var scores []int
	for i := 0; i < len(time); i = i + 1 {
		recordTime := time[i]
		recordDistance := distance[i]
		raceScore := evaluateRace(recordTime, recordDistance)
		scores = append(scores, raceScore)
	}

	return calcScoreProduct(scores)
}

func evaluateRace(recordTime int, recordDistance int) int {
	waysToWin := 0
	for waitTime := 0; waitTime < recordTime; waitTime = waitTime + 1 {
		calcDistance := waitTime * (recordTime - waitTime)
		if calcDistance > recordDistance {
			waysToWin = waysToWin + 1
		}
	}

	return waysToWin
}

func calcScoreProduct(scores []int) int {
	total := 1
	for _, score := range scores {
		total = total * score
	}

	return total
}

func parseInput(data []string) ([]int, []int, int, int) {
	var time, distance []int
	var compressedTime, compressedDistance int
	for i, line := range data {
		_, rightLine, _ := strings.Cut(line, ":")
		rawValues := strings.Split(strings.TrimSpace(rightLine), " ")
		for _, value := range rawValues {
			if value != "" {
				converted, _ := strconv.Atoi(value)
				if i == 0 {
					time = append(time, converted)
				} else {
					distance = append(distance, converted)
				}

			}
		}
		numberRegex := `\d`
		r, _ := regexp.Compile(numberRegex)
		digits := r.FindAllString(line, -1)
		compressed := strings.Join(digits, "")
		compressedNumber, _ := strconv.Atoi(compressed)
		if i == 0 {
			compressedTime = compressedNumber
		} else {
			compressedDistance = compressedNumber
		}
	}

	return time, distance, compressedTime, compressedDistance
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
