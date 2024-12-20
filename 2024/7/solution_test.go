package solution7

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example.txt")

	if result != 3749 {
		t.Fatalf(`Expected 3749, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 12553187650171 {
		t.Fatalf(`Expected 12553187650171, received %v`, result)
	}

}
