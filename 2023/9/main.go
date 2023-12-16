package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result1, result2 := solve("input.txt")

	log.Printf(`Result 1: %v`, result1)
	log.Printf(`Result 2: %v`, result2)
}

func solve(filename string) (int, int) {
	data := readInput(filename)
	sum := 0
	sumPrev := 0

	for _, line := range data {
		sensorData := parseLine(line)
		log.Println(sensorData)
		extrapolation := extrapolate(sensorData)
		downExtrapolation := extrapolateDown(sensorData)

		log.Printf(`next value ßby extrapolation: %v`, extrapolation)
		log.Printf(`prev value by extrapolation: %v`, downExtrapolation)
		sum = sum + extrapolation
		sumPrev = sumPrev + downExtrapolation
	}

	return sum, sumPrev
}

func extrapolateDown(data []int) int {
	var diff []int
	// build next diff array
	// calc diff between elements
	// if constant, extrapolate back
	// otherwise, call extrapolate recursively
	for i := 0; i < (len(data) - 1); i = i + 1 {
		diffVal := data[i+1] - data[i]
		diff = append(diff, diffVal)
	}

	isConstant, prevDiff := isConstantArray(diff)

	if !isConstant {
		prevDiff = extrapolateDown(diff)
	}

	nextValue := data[0] - prevDiff

	return nextValue

}

func extrapolate(data []int) int {
	var diff []int
	// build next diff array
	// calc diff between elements
	// if constant, extrapolate back
	// otherwise, call extrapolate recursively
	for i := 0; i < (len(data) - 1); i = i + 1 {
		diffVal := data[i+1] - data[i]
		diff = append(diff, diffVal)
	}

	isConstant, nextDiff := isConstantArray(diff)

	if !isConstant {
		nextDiff = extrapolate(diff)
	}

	nextValue := nextDiff + data[len(data)-1]

	return nextValue
}

func isConstantArray(data []int) (bool, int) {
	isConstant := true
	delta := data[0]
	for _, value := range data {
		if value != delta {
			isConstant = false
			break
		}
	}

	return isConstant, delta
}

func parseLine(line string) []int {
	var result []int
	rawSplits := strings.Split(line, " ")
	for _, strValue := range rawSplits {
		intValue, _ := strconv.Atoi(strValue)
		result = append(result, intValue)
	}

	return result
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
