package main

import "testing"

func Test_PartOne(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 4361},
		{filename: "sample_end.txt", want: 4361},
		{filename: "input_beginning.txt", want: 16846},
		{filename: "input_end.txt", want: 34092},
		{filename: "input_middle.txt", want: 38925},
		{filename: "input_middle_2.txt", want: 36330},
		{filename: "input.txt", want: 544664},
	}

	for _, tc := range tests {
		got := solvePartOne(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_PartTwo(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 467835},
		{filename: "input.txt", want: 84495585},
	}

	for _, tc := range tests {
		got := solvePartTwo(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
