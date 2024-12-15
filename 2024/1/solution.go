package solution1

import (
	"math"
	"scondon87/advent-ofcode/2024/readinput"
	"slices"
	"strconv"
	"strings"
)

func solution(filename string) int {
	var totalDiff float64
	input := readinput.Read(filename)

	a := []int{}
	b := []int{}

	for _, line := range input {
		splitLine := strings.Split(line, "   ")
		first, _ := strconv.Atoi(splitLine[0])
		second, _ := strconv.Atoi(splitLine[1])
		a = append(a, first)
		b = append(b, second)
	}

	slices.Sort(a)
	slices.Sort(b)

	for i := range a {
		diff := a[i] - b[i]
		absDiff := math.Abs(float64(diff))
		totalDiff += absDiff
	}

	return int(totalDiff)
}
