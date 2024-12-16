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

	matchGroups := findValidGroups(input[0])

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

		// fmt.Println(first, second)

		product := first * second
		sum += product
	}

	return sum
}

func findValidGroups(input string) []string {
	rawRegex := `mul\(\d{1,3},\d{1,3}\)`

	re, err := regexp.Compile(rawRegex)
	if err != nil {
		fmt.Println("ERROR: failed to compile regex")
	}

	return re.FindAllString(input, -1)
}

func findParameters(input string) ([]string, error) {
	rawRegex := `\d{1,3}`

	re, err := regexp.Compile(rawRegex)
	if err != nil {
		return nil, err
	}

	return re.FindAllString(input, -1), nil
}
