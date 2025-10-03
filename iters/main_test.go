package main

import (
	"iter"
	"reflect"
	"testing"
)

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
			got := []int{}
			iter := Range(test.from, test.to)
			for e := range iter {
				got = append(got, e)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
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
			name:      "No positive predicate",
			iter:      Range(1, 5),
			predicate: func(i int) bool { return false },
			want:      []int{},
		},
		{
			name:      "All positive predicate",
			iter:      Range(1, 5),
			predicate: func(i int) bool { return true },
			want:      []int{1, 2, 3, 4, 5},
		},
		{
			name:      "Even numbers only",
			iter:      Range(1, 5),
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{2, 4},
		},
		{
			name:      "Empy input, empty output",
			iter:      Range(6, 5),
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := []int{}
			for v := range Filter(test.iter, test.predicate) {
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
