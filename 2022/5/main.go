package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := solveChallenge("input.txt")

	log.Printf(`Result: %v`, result)
}

func solveChallenge(filename string) string {
	data, splitIndex := readInput(filename)
	layout, instructions := splitData(data, splitIndex)

	parsedLayout := parseLayout(layout)
	executedLayout := executeInstructions(parsedLayout, parseInstructions(instructions))

	return executedLayout.getMessage()
}

func executeInstructions(parsedLayout Layout, parsedInstructions []InstructionSet) Layout {
	for _, instruction := range parsedInstructions {
		parsedLayout.transfer(instruction)
	}

	return parsedLayout
}

func parseInstructions(instructions []string) []InstructionSet {
	var parsedInstructions []InstructionSet

	for _, rawInstruction := range instructions {
		parsedInstructions = append(parsedInstructions, buildInstructionSet(rawInstruction))
	}

	return parsedInstructions
}

func buildInstructionSet(raw string) InstructionSet {
	splitRaw := strings.Split(raw, " ")
	// index 1 is amount, 3 is source, 5 is destination
	amount, _ := strconv.Atoi(splitRaw[1])

	return InstructionSet{
		amount:      amount,
		source:      splitRaw[3],
		destination: splitRaw[5],
	}
}

func parseLayout(layout []string) Layout {
	numColumns := calcNumColumns(layout[len(layout)-1])
	parsedLayout := Layout{}
	parsedLayout = parsedLayout.init(numColumns)
	layout = layout[:len(layout)-1]

	return loadIntoParsedLayout(layout, parsedLayout)
}

func loadIntoParsedLayout(layout []string, parsedLayout Layout) Layout {
	// iterate in reverse
	for i := len(layout) - 1; i >= 0; i = i - 1 {
		splitLine := strings.Split(layout[i], "")
		for index, value := range splitLine {
			if value != " " && value != "[" && value != "]" {
				column := strconv.Itoa((index / 4) + 1)
				parsedLayout.push(column, value)
			}
		}
	}

	return parsedLayout
}

func calcNumColumns(line string) int {
	colNameDelimiter := "   "
	columnNames := strings.Split(line, colNameDelimiter)
	numColumns, _ := strconv.Atoi(columnNames[len(columnNames)-1])

	return numColumns
}

func splitData(data []string, splitIndex int) ([]string, []string) {
	layout := data[:splitIndex]
	instructions := data[splitIndex+1:]

	return layout, instructions
}

func readInput(filename string) ([]string, int) {
	splitIndex := -1
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
		if text == "" {
			splitIndex = len(dataSlice)
		}
		dataSlice = append(dataSlice, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataSlice, splitIndex
}

type Layout map[string][]string

func (l Layout) init(numColumns int) Layout {
	for i := 1; i <= numColumns; i = i + 1 {
		l[strconv.Itoa(i)] = []string{}
	}

	return l
}

func (l Layout) push(key string, value string) {
	l[key] = append(l[key], value)
}

func (l Layout) pop(key string) string {
	lastIndex := len(l[key]) - 1
	value := l[key][lastIndex]
	l[key] = l[key][:lastIndex]
	return value
}

func (l Layout) transfer(instruction InstructionSet) {
	for i := 0; i < instruction.amount; i = i + 1 {
		val := l.pop(instruction.source)
		l.push(instruction.destination, val)
	}
}

func (l Layout) getMessage() string {
	message := ""
	var keys []string
	for k := range l {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if len(l[key]) > 0 {
			message = message + l[key][len(l[key])-1]
		}
	}

	return message
}

type InstructionSet struct {
	amount      int
	source      string
	destination string
}
