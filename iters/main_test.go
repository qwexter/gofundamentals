package main

import (
	"iter"
	"reflect"
	"strconv"
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

func TestMap(t *testing.T) {
	// mapper is super primitive but it should care about all errors and nil value.
	// Otherwise it should be iter.Seq2 and work with errors somehow?
	mapper := func(v string) int {
		r, _ := strconv.Atoi(v)
		return r
	}

	tests := []struct {
		name string
		data []string
		want []int
	}{
		{
			name: "when converting empty iter then get an empy result iter",
			data: []string{},
			want: []int{},
		},
		{
			name: "when converting valid num string, then get a valid int slice",
			data: []string{"1", "2", "3"},
			want: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iter := Map(SliceAsIter(test.data), mapper)
			got := iterAsSlice(iter)
			assertEqual(t, test.name, got, test.want)
		})
	}
}

func TestTake(t *testing.T) {
	dataSeq := Range(0, 99)
	dataSlice := iterAsSlice(dataSeq)

	tests := []struct {
		name string
		take int
		want []int
	}{
		{
			name: "when take 0, then empty output",
			take: 0,
			want: []int{},
		},
		{
			name: "when take 1, then first element from input iterator",
			take: 1,
			want: dataSlice[:1],
		},
		{
			name: "when take 10, then iter of 10 first",
			take: 10,
			want: dataSlice[:10],
		},
		{
			name: "when take more then in input iterators, then return all available and thats all",
			take: 200,
			want: dataSlice,
		},
		{
			name: "when take is negative value, then empty output",
			take: -100,
			want: []int{},
		},
	}
	for _, test := range tests {
		iter := Take(dataSeq, test.take)
		got := iterAsSlice(iter)
		assertEqual(t, "when take is 1, then slice len is 1", got, test.want)
	}
}

func TestCycle(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		take  int
		want  []int
	}{
		{
			name:  "empty input, empty output",
			input: []int{},
			take:  10,
			want:  []int{},
		},
	}

	for _, test := range tests {
		iter := Cycle(test.input)
		wrap := Take(iter, test.take)
		got := iterAsSlice(wrap)
		assertEqual(t, test.name, got, test.want)
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
