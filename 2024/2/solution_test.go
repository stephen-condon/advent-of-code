package solution2

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 2 {
		t.Fatalf(`Expected 2, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 598 {
		t.Fatalf(`Expected 598, received %v`, result)
	}

}
