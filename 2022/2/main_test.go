package main

import "testing"

func TestPartOne_Sample(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 12},
		{filename: "input.txt", want: 11657},
	}

	for _, tc := range tests {
		got := solveChallenge(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
