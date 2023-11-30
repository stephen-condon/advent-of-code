package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	result := solvePartOne("input.txt")

	log.Printf(`Result: %v`, result)
}

func solvePartTwo(filename string) int {
	data := readInput(filename)
	buffer := strings.Split(data[0], "")
	index := processBuffer(buffer)

	return index
}

func solvePartOne(filename string) int {
	data := readInput(filename)
	buffer := strings.Split(data[0], "")
	index := processBuffer(buffer)

	return index
}

func processBuffer(buffer []string) int {
	// separate statements to ensure proper type for workingBuffer
	var workingBuffer BufferTest
	workingBuffer = buffer[:4]

	processedCharacters := -1
	// start iteration after preloading first 4 characters
	for i := 4; i < len(buffer); i = i + 1 {
		workingBuffer.rotate(buffer[i])
		if workingBuffer.test() {
			processedCharacters = i + 1
			break
		}
	}

	return processedCharacters
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

type BufferTest []string

func (b BufferTest) rotate(newData string) {
	oldBuffer := b
	for index := range oldBuffer {
		if index < 3 {
			b[index] = b[index+1]
		} else {
			// index == 3
			b[index] = newData
		}
	}
}

func (b BufferTest) test() bool {
	var workingBuffer []string
	isValid := true

	for _, value := range b {
		workingBuffer = append(workingBuffer, value)
	}

	sort.Strings(workingBuffer)

	for index := range workingBuffer {
		if index >= 1 {
			if workingBuffer[index] == workingBuffer[index-1] {
				isValid = false
				break
			}
		}
	}

	return isValid
}
