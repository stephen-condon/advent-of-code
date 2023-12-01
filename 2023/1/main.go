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
	sum := 0
	data := readInput(filename)

	r, _ := regexp.Compile(`\d`)

	for _, line := range data {
		matches := r.FindAllStringIndex(line, -1)

		firstDigit, _ := strconv.Atoi(line[matches[0][0]:matches[0][1]])
		secondDigit, _ := strconv.Atoi(line[matches[len(matches)-1][0]:matches[len(matches)-1][1]])

		calibrationValue := (firstDigit * 10) + secondDigit

		sum = sum + calibrationValue
	}

	return sum
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
