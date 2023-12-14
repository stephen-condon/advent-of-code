package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	result1 := solve("input.txt")

	log.Printf(`Result 1: %v`, result1)
}

func solve(filename string) int {
	data := readInput(filename)

	instructions, binaryTree := parseInput(data)

	currentPos := "AAA"
	numSteps := 0
	instructionIndex := 0

	for currentPos != "ZZZ" {
		if instructions[instructionIndex] == "L" {
			currentPos = binaryTree[currentPos][0]
		} else {
			currentPos = binaryTree[currentPos][1]
		}
		numSteps = numSteps + 1
		instructionIndex = instructionIndex + 1
		// reset instructionIndex if we get to the end
		if instructionIndex == len(instructions) {
			instructionIndex = 0
		}
	}

	return numSteps
}

func parseInput(data []string) ([]string, Tree) {
	binaryTree := Tree{}
	var instructions []string

	inputRegex := `[A-Z]{3}`

	r, _ := regexp.Compile(inputRegex)

	instructions = strings.Split(data[0], "")

	treeData := data[2:]

	for _, line := range treeData {
		matches := r.FindAllString(line, -1)
		children := [2]string{matches[1], matches[2]}
		binaryTree[matches[0]] = children
	}

	return instructions, binaryTree
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

type Tree map[string][2]string
