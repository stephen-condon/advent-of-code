package main

import "testing"

func TestPartOne_Sample(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 70},
		{filename: "input.txt", want: 2752},
	}

	for _, tc := range tests {
		got := solveChallenge(tc.filename)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestPartOne_InitializePriorities(t *testing.T) {
	expectedLength := 52
	expectedUpperA := 27
	expectedLowerA := 1
	initializePriorities()
	got := priorityMap
	if len(got) != expectedLength {
		t.Fatalf("expected length: %v, got: %v", expectedLength, got)
	}
	if got["A"] != expectedUpperA {
		t.Fatalf("expected length: %v, got: %v", expectedUpperA, got)
	}
	if got["a"] != expectedLowerA {
		t.Fatalf("expected length: %v, got: %v", expectedLowerA, got)
	}
}
