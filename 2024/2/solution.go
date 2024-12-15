package solution2

import (
	"math"
	"scondon87/advent-ofcode/2024/readinput"
	"strconv"
	"strings"
)

func solution(filename string) int {
	count := 0
	input := readinput.Read(filename)

	for _, report := range input {
		rawReport := strings.Split(report, " ")
		intReport := []int{}
		for _, value := range rawReport {
			intValue, _ := strconv.Atoi(value)
			intReport = append(intReport, intValue)
		}
		safe := isSafe(intReport)
		if safe {
			count++
		}
	}

	return count
}

func isSafe(report []int) bool {
	// get allowed trend from first two values
	// calculate as we go, if it ever flips from what it's set at, mark unsafe, early exit
	trend := checkTrend(report[0], report[1])

	for i := range report[:len(report)-1] {
		if trend != checkTrend(report[i], report[i+1]) {
			return false
		}

		if !checkStepSize(report[i], report[i+1]) {
			return false
		}
	}

	return true
}

func checkTrend(a int, b int) int {
	trend := 0
	if a > b {
		trend = -1
	} else if a < b {
		trend = 1
	}

	return trend
}

func checkStepSize(a int, b int) bool {
	diff := a - b
	absDiff := math.Abs(float64(diff))
	if absDiff < 1 || absDiff > 3 {
		return false
	}

	return true
}
