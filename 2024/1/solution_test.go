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
