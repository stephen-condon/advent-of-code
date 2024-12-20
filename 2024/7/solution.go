package solution7

import (
	"fmt"
	"scondon87/advent-ofcode/2024/readinput"
	"strconv"
	"strings"
)

type equation struct {
	answer     int
	components []int
}

func (e *equation) evalPermutations(operators []string) int {
	result := 0

	// Get all possible combinations of operators for equation size
	numOperators := len(e.components) - 1
	operatorCombos := make([][]string, 0)

	// Generate all possible operator combinations
	var generateCombos func(combo []string)
	generateCombos = func(combo []string) {
		if len(combo) == numOperators {
			operatorCombos = append(operatorCombos, append([]string{}, combo...))
			return
		}
		for _, op := range operators {
			generateCombos(append(combo, op))
		}
	}
	generateCombos([]string{})

	// For each combination, evaluate the equation
	for _, combo := range operatorCombos {
		total := e.components[0]
		for i := 0; i < len(combo); i++ {
			if combo[i] == "+" {
				total += e.components[i+1]
			} else if combo[i] == "*" {
				total *= e.components[i+1]
			}
		}
		if total == e.answer {
			result = e.answer
			break
		}
	}

	return result
}

func (e *equation) evalPermutations2(operators []string) int {
	result := 0

	// Get all possible combinations of operators for equation size
	numOperators := len(e.components) - 1
	operatorCombos := make([][]string, 0)

	// Generate all possible operator combinations
	var generateCombos func(combo []string)
	generateCombos = func(combo []string) {
		if len(combo) == numOperators {
			operatorCombos = append(operatorCombos, append([]string{}, combo...))
			return
		}
		for _, op := range operators {
			generateCombos(append(combo, op))
		}
	}
	generateCombos([]string{})

	// For each combination, evaluate the equation
	for _, combo := range operatorCombos {
		total := e.components[0]
		for i := 0; i < len(combo); i++ {
			if combo[i] == "+" {
				total += e.components[i+1]
			} else if combo[i] == "*" {
				total *= e.components[i+1]
			} else if combo[i] == "||" {
				// Convert current total and next number to strings and concatenate
				totalStr := strconv.Itoa(total)
				nextStr := strconv.Itoa(e.components[i+1])
				concatenated, _ := strconv.Atoi(totalStr + nextStr)
				total = concatenated
			}
		}
		if total == e.answer {
			result = e.answer
			break
		}
	}

	return result
}

func solution(filename string) int {
	sum := 0
	input := readinput.Read(filename)

	operators := []string{"+", "*"}

	for _, line := range input {
		splitEq := strings.Split(line, ": ")
		eqAnswer, err := strconv.Atoi(splitEq[0])
		if err != nil {
			fmt.Println(err)
		}
		eqComponents := strings.Split(splitEq[1], " ")
		theComponents := []int{}
		for _, comp := range eqComponents {
			theComp, err := strconv.Atoi(comp)
			if err != nil {
				fmt.Println(err)
			}
			theComponents = append(theComponents, theComp)
		}
		theEquation := equation{
			answer:     eqAnswer,
			components: theComponents,
		}

		result := theEquation.evalPermutations(operators)

		sum += result
	}

	return sum
}

func solution2(filename string) int {
	sum := 0
	input := readinput.Read(filename)

	operators := []string{"+", "*", "||"}

	for _, line := range input {
		splitEq := strings.Split(line, ": ")
		eqAnswer, err := strconv.Atoi(splitEq[0])
		if err != nil {
			fmt.Println(err)
		}
		eqComponents := strings.Split(splitEq[1], " ")
		theComponents := []int{}
		for _, comp := range eqComponents {
			theComp, err := strconv.Atoi(comp)
			if err != nil {
				fmt.Println(err)
			}
			theComponents = append(theComponents, theComp)
		}
		theEquation := equation{
			answer:     eqAnswer,
			components: theComponents,
		}

		result := theEquation.evalPermutations2(operators)

		sum += result
	}

	return sum
}
