package main

import "testing"

func Test_PartOne(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 142},
		{filename: "input.txt", want: 54081},
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
		{filename: "sample2.txt", want: 281},
		{filename: "bug.txt", want: 12},
		{filename: "input.txt", want: 54649},
	}

	for _, tc := range tests {
		got := solvePartTwo(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
