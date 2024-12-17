package solution5

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 143 {
		t.Fatalf(`Expected 143, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 4766 {
		t.Fatalf(`Expected 4766, received %v`, result)
	}

}
