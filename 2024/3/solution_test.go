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
