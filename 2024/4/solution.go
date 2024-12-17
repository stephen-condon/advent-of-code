package solution4

import (
	"fmt"
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
		fmt.Println(split)
	}

	// loop through by row
	for i := range data {
		xIndices := searchTerm.findX(data[i])

		for _, x := range xIndices {
			fmt.Println("precheck", i, x)
			found := searchTerm.findWord(x, i, data)
			fmt.Println("found", i, x, found)
			count += found
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
		fmt.Println("up", y, x)
		count++
	}
	if y < len(data)-3 && data[y+1][x] == wf[1] && data[y+2][x] == wf[2] && data[y+3][x] == wf[3] {
		// check down
		fmt.Println("down", y, x)
		count++
	}
	if x >= 3 && data[y][x-1] == wf[1] && data[y][x-2] == wf[2] && data[y][x-3] == wf[3] {
		// check left
		fmt.Println("left", y, x)
		count++
	}
	if x < len(data[0])-3 && data[y][x+1] == wf[1] && data[y][x+2] == wf[2] && data[y][x+3] == wf[3] {
		// check right
		fmt.Println("right", y, x)
		count++
	}
	if y >= 3 && x < len(data[0])-3 && data[y-1][x+1] == wf[1] && data[y-2][x+2] == wf[2] && data[y-3][x+3] == wf[3] {
		// check up right
		fmt.Println("up-right", y, x)
		count++
	}
	if y >= 3 && x >= 3 && data[y-1][x-1] == wf[1] && data[y-2][x-2] == wf[2] && data[y-3][x-3] == wf[3] {
		// check up left
		fmt.Println("up-left", y, x)
		count++
	}
	if y < len(data)-3 && x < len(data[0])-3 && data[y+1][x+1] == wf[1] && data[y+2][x+2] == wf[2] && data[y+3][x+3] == wf[3] {
		// check down right
		fmt.Println("down-right", y, x)
		count++
	}
	if y < len(data)-3 && x >= 3 && data[y+1][x-1] == wf[1] && data[y+2][x-2] == wf[2] && data[y+3][x-3] == wf[3] {
		// check down left
		fmt.Println("down-left", y, x)
		count++
	}

	return count
}
