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
