package main

import (
	"iter"
	"reflect"
	"testing"
)

func TestIterAsSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "when empty slice then empty iter",
			input: []int{},
		},
		{
			name:  "when non empty slice then same values as iter",
			input: []int{1, 3, 2, 4},
		},
	}

	for _, test := range tests {
		iter := SliceAsIter(test.input)
		got := iterAsSlice(iter)
		assertEqual(t, test.name, got, test.input)
	}
}

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		from, to int
		want     []int
	}{
		{
			name: "range from 0 to 5 including",
			from: 0,
			to:   5,
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "range from 3 to 5 including",
			from: 3,
			to:   5,
			want: []int{3, 4, 5},
		},
		{
			name: "range from 5 to 4 including",
			from: 5,
			to:   4,
			want: []int{},
		},
		{
			name: "range from -5 to -4 including",
			from: -5,
			to:   -4,
			want: []int{-5, -4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := iterAsSlice(Range(test.from, test.to))
			assertEqual(t, test.name, got, test.want)
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		iter      iter.Seq[int]
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "no positive predicate",
			iter:      Range(1, 5),
			predicate: func(i int) bool { return false },
			want:      []int{},
		},
		{
			name:      "all positive predicate",
			iter:      Range(1, 5),
			predicate: func(i int) bool { return true },
			want:      []int{1, 2, 3, 4, 5},
		},
		{
			name:      "even numbers only",
			iter:      Range(1, 5),
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{2, 4},
		},
		{
			name:      "empy input, empty output",
			iter:      Range(6, 5),
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := iterAsSlice(Filter(test.iter, test.predicate))
			assertEqual(t, test.name, got, test.want)
		})
	}
}

func iterAsSlice[T any](iter iter.Seq[T]) []T {
	got := []T{}
	for v := range iter {
		got = append(got, v)
	}
	return got
}

func assertEqual[T any](t *testing.T, test string, got, want []T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("for test: %s error, got %v, but want %v", test, got, want)
	}
}
