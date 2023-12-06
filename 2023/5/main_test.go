package main

import (
	"testing"
)

func Test_PartOne(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 35},
		// {filename: "input.txt", want: 26443},
	}

	for _, tc := range tests {
		part1, _ := solve(tc.filename, true)
		if part1 != tc.want {
			t.Fatalf("expected: %v, part1: %v", tc.want, part1)
		}
	}
}
