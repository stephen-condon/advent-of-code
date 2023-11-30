package main

import "testing"

func Test_PartOne_Sample(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample5.txt", want: 5},
		{filename: "sample6.txt", want: 6},
		{filename: "sample7.txt", want: 7},
		{filename: "sample10.txt", want: 10},
		{filename: "sample11.txt", want: 11},
		{filename: "input.txt", want: 1566},
	}

	for _, tc := range tests {
		got := solvePartOne(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_PartTwo_Sample(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample5.txt", want: 23},
		{filename: "sample6.txt", want: 23},
		{filename: "sample7.txt", want: 19},
		{filename: "sample10.txt", want: 29},
		{filename: "sample11.txt", want: 26},
		{filename: "input.txt", want: 2265},
	}

	for _, tc := range tests {
		got := solvePartTwo(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
