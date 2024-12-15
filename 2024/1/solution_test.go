package solution1

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 11 {
		t.Fatalf(`Expected 11, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 1938424 {
		t.Fatalf(`Expected 1938424, received %v`, result)
	}

}

func TestSolution2Example(t *testing.T) {
	result := solution2("./example.txt")

	if result != 31 {
		t.Fatalf(`Expected 31, received %v`, result)
	}

}

func TestSolution2Input(t *testing.T) {
	result := solution2("./input.txt")

	if result != 22014209 {
		t.Fatalf(`Expected 22014209, received %v`, result)
	}

}
