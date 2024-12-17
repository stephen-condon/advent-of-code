package solution4

import (
	"scondon87/advent-ofcode/2024/readinput"
	"strings"
)

type wordFind []string

func solution(filename string) int {
	count := 0
	input := readinput.Read(filename)
	var data [][]string = make([][]string, len(input))

	searchTerm := wordFind{"X", "M", "A", "S"}

	// populate data as [][]string
	for i := range input {
		split := strings.Split(input[i], "")
		data[i] = split
	}

	// loop through by row
	for i := range data {
		xIndices := searchTerm.findX(data[i])

		for _, x := range xIndices {
			found := searchTerm.findWord(x, i, data)
			count += found
		}
	}

	return count
}

func solution2(filename string) int {
	count := 0
	input := readinput.Read(filename)
	var data [][]string = make([][]string, len(input))

	searchTerm := wordFind{"M", "A", "S"}

	// populate data as [][]string
	for i := range input {
		split := strings.Split(input[i], "")
		data[i] = split
	}

	// find "A", which is the center, then need to check diagonals, opposites need to be M/S
	// loop through by row
	for i := range data {
		if i != 0 && i != len(data)-1 {
			aIndices := searchTerm.findA(data[i], len(data[i]))

			for _, x := range aIndices {
				found := searchTerm.findPattern(x, i, data)
				count += found
			}
		}
	}

	return count

}

func (wf wordFind) findX(input []string) []int {
	result := []int{}
	for i := range input {
		if input[i] == wf[0] {
			result = append(result, i)
		}
	}

	return result
}

func (wf wordFind) findWord(x int, y int, data [][]string) int {
	count := 0
	if y >= 3 && data[y-1][x] == wf[1] && data[y-2][x] == wf[2] && data[y-3][x] == wf[3] {
		// check up
		count++
	}
	if y < len(data)-3 && data[y+1][x] == wf[1] && data[y+2][x] == wf[2] && data[y+3][x] == wf[3] {
		// check down
		count++
	}
	if x >= 3 && data[y][x-1] == wf[1] && data[y][x-2] == wf[2] && data[y][x-3] == wf[3] {
		// check left
		count++
	}
	if x < len(data[0])-3 && data[y][x+1] == wf[1] && data[y][x+2] == wf[2] && data[y][x+3] == wf[3] {
		// check right
		count++
	}
	if y >= 3 && x < len(data[0])-3 && data[y-1][x+1] == wf[1] && data[y-2][x+2] == wf[2] && data[y-3][x+3] == wf[3] {
		// check up right
		count++
	}
	if y >= 3 && x >= 3 && data[y-1][x-1] == wf[1] && data[y-2][x-2] == wf[2] && data[y-3][x-3] == wf[3] {
		// check up left
		count++
	}
	if y < len(data)-3 && x < len(data[0])-3 && data[y+1][x+1] == wf[1] && data[y+2][x+2] == wf[2] && data[y+3][x+3] == wf[3] {
		// check down right
		count++
	}
	if y < len(data)-3 && x >= 3 && data[y+1][x-1] == wf[1] && data[y+2][x-2] == wf[2] && data[y+3][x-3] == wf[3] {
		// check down left
		count++
	}

	return count
}

func (wf wordFind) findA(input []string, length int) []int {
	result := []int{}
	for i := range input {
		if input[i] == wf[1] && i != 0 && i != length-1 {
			result = append(result, i)
		}
	}

	return result
}

func (wf wordFind) findPattern(x int, y int, data [][]string) int {
	count := 0
	if data[y-1][x-1] == "M" && data[y+1][x+1] == "S" {
		if data[y-1][x+1] == "M" && data[y+1][x-1] == "S" {
			count++
		} else if data[y-1][x+1] == "S" && data[y+1][x-1] == "M" {
			count++
		}
	} else if data[y-1][x-1] == "S" && data[y+1][x+1] == "M" {
		if data[y-1][x+1] == "M" && data[y+1][x-1] == "S" {
			count++
		} else if data[y-1][x+1] == "S" && data[y+1][x-1] == "M" {
			count++
		}
	}

	return count
}
