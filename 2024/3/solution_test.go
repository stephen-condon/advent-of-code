package solution3

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 161 {
		t.Fatalf(`Expected 161, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 175615763 {
		t.Fatalf(`Expected 175615763, received %v`, result)
	}

}

func TestSolution2Example(t *testing.T) {
	result := solution2("./example2.txt")

	if result != 48 {
		t.Fatalf(`Expected 48, received %v`, result)
	}

}

func TestSolution2Example3(t *testing.T) {
	result := solution2("./example3.txt")

	if result != 2791426 {
		t.Fatalf(`Expected 2791426, received %v`, result)
	}

}

func TestSolution2Input(t *testing.T) {
	result := solution2("./input.txt")

	if result != 74361272 {
		t.Fatalf(`Expected 74361272, received %v`, result)
	}

}
