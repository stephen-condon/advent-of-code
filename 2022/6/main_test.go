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
