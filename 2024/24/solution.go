package solution24

import (
	"fmt"
	"scondon87/advent-ofcode/2024/readinput"
	"sort"
	"strconv"
	"strings"
)

type gateData struct {
	first     string
	second    string
	operation string
	output    string
}

type ValueStore map[string]int

func solution(filename string) int {
	input := readinput.Read(filename)
	values, gates := parse(input)

	for len(gates) > 0 {
		for i, gate := range gates {
			ready := gate.hasAllValues(values)
			if ready {
				values[gate.output] = gate.calc(values)
				if i == len(gates)-1 {
					gates = append(gates[:i], gates[i+1:]...)
				}
			}
		}
	}

	values = values.startsWithZ()
	return values.collate()
}

func parse(input []string) (ValueStore, []gateData) {
	initialValues := make(ValueStore)
	gates := []gateData{}
	donePartOne := false

	for _, line := range input {
		if !donePartOne {
			if line == "" {
				donePartOne = true
			} else {
				splitLine := strings.Split(line, ": ")
				val, err := strconv.Atoi(splitLine[1])
				if err != nil {
					return nil, nil
				}
				initialValues[splitLine[0]] = val
			}
		} else {
			new := newGate(line)
			gates = append(gates, new)
		}
	}

	return initialValues, gates

}

func newGate(line string) gateData {
	splitLine := strings.Split(line, " ")
	return gateData{
		first:     splitLine[0],
		second:    splitLine[2],
		operation: splitLine[1],
		output:    splitLine[4],
	}
}

func (gd *gateData) hasAllValues(values ValueStore) bool {
	if _, ok := values[gd.first]; !ok {
		return false
	}
	if _, ok := values[gd.second]; !ok {
		return false
	}
	return true
}

func (gd *gateData) calc(values ValueStore) int {
	if gd.operation == "AND" {
		if values[gd.first] == 1 && values[gd.second] == 1 {
			return 1
		}
	} else if gd.operation == "OR" {
		if values[gd.first] == 1 || values[gd.second] == 1 {
			return 1
		}
	} else if gd.operation == "XOR" {
		if (values[gd.first] == 1 && values[gd.second] == 0) || (values[gd.first] == 0 && values[gd.second] == 1) {
			return 1
		}
	}

	return 0
}

func (vs *ValueStore) startsWithZ() ValueStore {
	new := make(ValueStore)
	for key, val := range *vs {
		if key[0] == 'z' {
			new[key] = val
		}
	}

	return new
}

func (vs *ValueStore) collate() int {
	// combine the values in the value store as binary and convert to integer
	store := *vs
	binary := ""

	keys := make([]string, 0)
	for k := range store {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		binary = fmt.Sprintf("%v%v", strconv.Itoa(store[k]), binary)
	}

	splitBinary := strings.Split(binary, "")

	sum := 0
	for i := 0; i < len(splitBinary); i++ {
		val, _ := strconv.Atoi(splitBinary[i])
		sum = (sum * 2) + val
	}

	return sum
}
