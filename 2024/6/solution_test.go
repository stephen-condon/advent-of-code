package solution6

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 41 {
		t.Fatalf(`Expected 41, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 4433 {
		t.Fatalf(`Expected 4433, received %v`, result)
	}

}
