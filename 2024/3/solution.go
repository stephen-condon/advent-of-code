package solution3

import (
	"fmt"
	"regexp"
	"scondon87/advent-ofcode/2024/readinput"
	"strconv"
)

func solution(filename string) int {
	sum := 0
	input := readinput.Read(filename)

	matchGroups, _ := findValidGroups(input[0])

	for _, match := range matchGroups {
		stringParams, err := findParameters(match)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		first, err := strconv.Atoi(stringParams[0])
		if err != nil {
			fmt.Println(err)
			return -1
		}
		second, err := strconv.Atoi(stringParams[1])
		if err != nil {
			fmt.Println(err)
			return -1
		}

		product := first * second
		sum += product
	}

	return sum
}

func solution2(filename string) int {
	process := true
	sum := 0
	input := readinput.Read(filename)

	groupIndices, _ := findGroupIndices(input[0])
	doIndices, _ := findDoIndices(input[0])
	dontIndices, _ := findDontIndices(input[0])

	// check next do & don't, if first index < either of them, process mul using current state

	for len(groupIndices) > 0 {
		nextDo := -1
		nextDont := -1
		if len(doIndices) > 0 {
			nextDo = doIndices[0][0]
		}
		if len(dontIndices) > 0 {
			nextDont = dontIndices[0][0]
		}

		nextGroup := groupIndices[0]

		if nextDo < nextGroup[0] && nextDo != -1 {
			process = true
			doIndices = doIndices[1:]
		} else if nextDont < nextGroup[0] && nextDont != -1 {
			process = false
			dontIndices = dontIndices[1:]
		} else {
			if process {
				// process group
				stringParams, err := findParameters(input[0][nextGroup[0]:nextGroup[1]])
				if err != nil {
					fmt.Println(err)
					return -1
				}
				first, err := strconv.Atoi(stringParams[0])
				if err != nil {
					fmt.Println(err)
					return -1
				}
				second, err := strconv.Atoi(stringParams[1])
				if err != nil {
					fmt.Println(err)
					return -1
				}

				product := first * second
				sum += product
			}
			groupIndices = groupIndices[1:]
		}

	}

	return sum
}

func findDoIndices(input string) ([][]int, error) {
	rawRegex := `do\(\)`

	matches, err := findIndexRegex(input, rawRegex)

	return matches, err
}

func findDontIndices(input string) ([][]int, error) {
	rawRegex := `don\'t\(\)`

	matches, err := findIndexRegex(input, rawRegex)

	return matches, err
}

func findGroupIndices(input string) ([][]int, error) {
	rawRegex := `mul\(\d{1,3},\d{1,3}\)`

	matches, err := findIndexRegex(input, rawRegex)

	return matches, err
}

func findValidGroups(input string) ([]string, error) {
	rawRegex := `mul\(\d{1,3},\d{1,3}\)`

	matches, err := findStringRegex(input, rawRegex)

	return matches, err
}

func findParameters(input string) ([]string, error) {
	rawRegex := `\d{1,3}`

	matches, err := findStringRegex(input, rawRegex)

	return matches, err
}

func findStringRegex(input string, regex string) ([]string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("ERROR: failed to compile regex")
		return nil, err
	}

	return re.FindAllString(input, -1), nil
}

func findIndexRegex(input string, regex string) ([][]int, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("ERROR: failed to compile regex")
		return nil, err
	}

	return re.FindAllStringIndex(input, -1), nil
}
