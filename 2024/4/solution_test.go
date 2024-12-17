package solution4

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 18 {
		t.Fatalf(`Expected 18, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 2562 {
		t.Fatalf(`Expected 2562, received %v`, result)
	}

}

func TestSolution2Example(t *testing.T) {
	result := solution2("./example.txt")

	if result != 9 {
		t.Fatalf(`Expected 9, received %v`, result)
	}

}

func TestSolution2Input(t *testing.T) {
	result := solution2("./input.txt")

	if result != 1902 {
		t.Fatalf(`Expected 1902, received %v`, result)
	}

}
