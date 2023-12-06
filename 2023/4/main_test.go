package main

import (
	"math"
	"reflect"
	"testing"
)

func Test_PartOne(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{filename: "sample.txt", want: 13},
		{filename: "input.txt", want: 26443},
	}

	for _, tc := range tests {
		got := solve(tc.filename, true)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_CalculateScore(t *testing.T) {
	type test struct {
		count int
		want  int
	}

	tests := []test{
		{count: 0, want: 0},
		{count: 1, want: 1},
		{count: 2, want: 2},
		{count: 3, want: 4},
		{count: 4, want: 8},
		{count: 5, want: 16},
		{count: 6, want: 32},
		{count: 7, want: 64},
		{count: 8, want: 128},
		{count: 9, want: 256},
		{count: -1, want: 0},
		{count: math.MinInt, want: 0},
	}

	for _, tc := range tests {
		got := calculateScore(tc.count)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_TrimString(t *testing.T) {
	type test struct {
		str  string
		want string
	}

	tests := []test{
		{str: "abc", want: "abc"},
		{str: " abc", want: "abc"},
		{str: "abc ", want: "abc"},
		{str: " abc ", want: "abc"},
		{str: "a b c", want: "a b c"},
	}

	for _, tc := range tests {
		got := trimString(tc.str)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_ConvertStringArrayToInts(t *testing.T) {
	type test struct {
		arr  []string
		want []int
	}

	tests := []test{
		{arr: []string{"3", "7"}, want: []int{3, 7}},
		{arr: []string{"3", " ", "7"}, want: []int{3, 7}},
	}

	for _, tc := range tests {
		got := convertStringArrayToInts(tc.arr)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_CountMatches(t *testing.T) {
	type test struct {
		winners []int
		ours    []int
		want    int
	}

	tests := []test{
		{winners: []int{1, 2}, ours: []int{1, 2}, want: 2},
		{winners: []int{1, 2}, ours: []int{3, 4}, want: 0},
		{winners: []int{1, 2}, ours: []int{3, 2}, want: 1},
		{winners: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, ours: []int{5, 9}, want: 2},
	}

	for _, tc := range tests {
		got := countMatches(tc.winners, tc.ours)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
