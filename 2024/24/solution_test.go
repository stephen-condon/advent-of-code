package solution24

import "testing"

func TestSolutionExample(t *testing.T) {
	result := solution("./example1.txt")

	if result != 4 {
		t.Fatalf(`Expected 4, received %v`, result)
	}

}

func TestSolutionExample2(t *testing.T) {
	result := solution("./example2.txt")

	if result != 2024 {
		t.Fatalf(`Expected 2024, received %v`, result)
	}

}

func TestSolutionInput(t *testing.T) {
	result := solution("./input.txt")

	if result != 58639252480880 {
		t.Fatalf(`Expected 58639252480880, received %v`, result)
	}

}

func TestCollate1(t *testing.T) {
	initial := ValueStore{
		"z00": 0,
		"z01": 0,
		"z02": 1,
	}

	expected := 4

	result := initial.collate()

	if result != expected {
		t.Fatalf(`Expected %v, received %v`, expected, result)
	}

}

func TestCollate2(t *testing.T) {
	initial := ValueStore{
		"z00": 0,
		"z01": 0,
		"z02": 0,
		"z03": 1,
		"z04": 0,
		"z05": 1,
		"z06": 1,
		"z07": 1,
		"z08": 1,
		"z09": 1,
		"z10": 1,
		"z11": 0,
		"z12": 0,
	}

	// 0011111101000

	expected := 2024

	result := initial.collate()

	if result != expected {
		t.Fatalf(`Expected %v, received %v`, expected, result)
	}

}
