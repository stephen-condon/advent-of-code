package solution25

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 3 {
		t.Fatalf(`Expected 3, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 2815 {
		t.Fatalf(`Expected 2815, received %v`, result)
	}

}
