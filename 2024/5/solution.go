package solution5

import (
	"fmt"
	"regexp"
	"scondon87/advent-ofcode/2024/readinput"
	"strconv"
	"strings"
)

type ruleObj []string

func solution(filename string) int {
	result := 0
	input := readinput.Read(filename)

	reachedDivider := false
	ruleSets := []ruleObj{}
	pageSets := []string{}
	correctPageSets := []string{}

	for _, line := range input {
		if line == "" {
			reachedDivider = true
		} else if reachedDivider {
			// pagesets
			pageSets = append(pageSets, line)
		} else {
			// rules
			new := strings.Split(line, "|")
			ruleSets = append(ruleSets, new)
		}
	}

	fmt.Println(ruleSets)
	fmt.Println(pageSets)
	for _, page := range pageSets {
		successful := true
		for _, rule := range ruleSets {
			result := rule.check(page)
			if result == -2 {
				// regex compile error
				return -2
			} else if result == -1 {
				// rule doesn't apply
			} else if result == 0 {
				// applied unsuccessfully
				successful = false
				break
			} else if result == 1 {
				// applied successfully
			} else {
				// should not happen
				return -99
			}
		}
		if successful {
			correctPageSets = append(correctPageSets, page)
		}
	}

	for _, page := range correctPageSets {
		split := strings.Split(page, ",")
		middle := (len(split) - 1) / 2
		middleValue, err := strconv.Atoi(split[middle])
		if err != nil {
			return -3
		}
		result += middleValue
	}

	return result
}

func (r ruleObj) check(input string) int {
	// check each individual number's presence
	orderRegex := ``

	for i, param := range r {
		rawRegex := fmt.Sprintf(`%v`, param)
		if i == 0 {
			orderRegex = fmt.Sprintf(`%v(%v)`, orderRegex, param)
		} else {
			orderRegex = fmt.Sprintf(`%v.*(%v)`, orderRegex, param)
		}
		re, err := regexp.Compile(rawRegex)
		if err != nil {
			fmt.Println(err)
			return -2
		}
		indices := re.FindAllStringIndex(input, -1)
		if len(indices) == 0 {
			return -1
		}
	}

	// if present, check order
	re, err := regexp.Compile(orderRegex)
	if err != nil {
		fmt.Println(err)
		return -2
	}
	indices := re.FindAllStringIndex(input, -1)
	if len(indices) == 0 {
		return 0
	}

	return 1
}

// need to apply the individual regexes to confirm all pages are present for the rule to apply, if it is, apply the ordering check
// (n1).*(n2)
