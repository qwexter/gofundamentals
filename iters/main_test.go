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
		{
			name: "range from 0 to 0, empty slice",
			from: 0,
			to:   0,
			want: []int{},
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
		assertEqual(t, test.name, got, test.want)
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
		{
			name:  "when we have take less than clice to cycle, then we get only taken count",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			take:  3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "when slice to cycle is shorter than taken count, then we get repeated slice bounded by taken",
			input: []int{0, 1},
			take:  7,
			want:  []int{0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, test := range tests {
		iter := Cycle(test.input)
		wrap := Take(iter, test.take)
		got := iterAsSlice(wrap)
		assertEqual(t, test.name, got, test.want)
	}
}

func TestEnumerate(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []pair[int, string]
	}{
		{
			name: "when empty input, then empty output",
			in:   []string{},
			want: []pair[int, string]{},
		},
		{
			name: "when non-empty input, then enumerated output",
			in:   []string{"0", "1", "2"},
			want: []pair[int, string]{{0, "0"}, {1, "1"}, {2, "2"}},
		},
	}

	for _, test := range tests {
		iter := Enumerate(test.in)

		got := make([]pair[int, string], len(test.in))
		for i, v := range iter {
			got[i] = pair[int, string]{i, v}
		}
		assertEqual(t, "when get input then get enumerated output", got, test.want)
	}
}

func TestZip(t *testing.T) {
	type TZip[T, U any] struct {
		name string
		in1  []T
		in2  []U
		want []pair[T, U]
	}
	tests := []TZip[int, string]{
		{
			name: "when both empty, then empty seq",
			in1:  []int{},
			in2:  []string{},
			want: []pair[int, string]{},
		},
		{
			name: "when left is shorter, then result is same length as shortest",
			in1:  []int{0, 1, 2},
			in2:  []string{"0", "1", "2", "3", "4"},
			want: []pair[int, string]{
				{first: 0, second: "0"},
				{first: 1, second: "1"},
				{first: 2, second: "2"},
			},
		},
		{
			name: "when right is shorter, then result is same length as shortest",
			in1:  []int{0, 1, 2, 3, 4},
			in2:  []string{"0", "1"},
			want: []pair[int, string]{
				{first: 0, second: "0"},
				{first: 1, second: "1"},
			},
		},
		{
			name: "when both same non-empty length, then same length seq",
			in1:  []int{0, 1, 2, 3, 4},
			in2:  []string{"0", "1", "2", "3", "4"},
			want: []pair[int, string]{
				{first: 0, second: "0"},
				{first: 1, second: "1"},
				{first: 2, second: "2"},
				{first: 3, second: "3"},
				{first: 4, second: "4"},
			},
		},
	}

	for _, test := range tests {
		seq1 := SliceAsIter(test.in1)
		seq2 := SliceAsIter(test.in2)

		got := iter2AsSlice(Zip(seq1, seq2))
		assertEqual(t, test.name, got, test.want)
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want []int
	}{
		{
			"when empty input, get an empty output",
			[][]int{},
			[]int{},
		},
		{
			"when input of empty iters, then empty output",
			[][]int{{}, {}, {}},
			[]int{},
		},
		{
			"when input mixed empty and non-empty, then output only non-empy values",
			[][]int{{}, {1, 2, 3}, {}},
			[]int{1, 2, 3},
		},
		{
			"when non empty input, then output only non-empy values",
			[][]int{{0}, {1, 2, 3}, {4}},
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		in := SliceAsIter(test.in)
		got := iterAsSlice(Flatten(in))
		assertEqual(t, test.name, got, test.want)
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name    string
		in      pair[int, int]
		chunkBy int
		want    [][]int
	}{
		{
			name:    "when empty input and zero chunkBy, then empty output",
			in:      pair[int, int]{0, 0},
			chunkBy: 0,
			want:    [][]int{},
		},
		{
			name:    "when input is zero length and chunk > 0, then empty output",
			in:      pair[int, int]{0, 0},
			chunkBy: 0,
			want:    [][]int{},
		},
		{
			name:    "when input non empty and chunk is 0, then empty output",
			in:      pair[int, int]{0, 5},
			chunkBy: 0,
			want:    [][]int{},
		},
		{
			name:    "when input non empty and chunk less then input length, then chunked output",
			in:      pair[int, int]{0, 5},
			chunkBy: 1,
			want:    [][]int{{0}, {1}, {2}, {3}, {4}, {5}},
		},
		{
			name:    "when input non empty and chunk more then input length, then one chunk as an output",
			in:      pair[int, int]{0, 5},
			chunkBy: 6,
			want:    [][]int{{0, 1, 2, 3, 4, 5}},
		},
	}

	for _, test := range tests {
		in := Range(test.in.first, test.in.second)
		got := iterAsSlice(Chunk(in, test.chunkBy))
		assertEqual(t, test.name, got, test.want)
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		in      iter.Seq[int]
		initial int
		want    int
	}{
		{
			in:      Range(1, 5),
			initial: 0,
			want:    15,
		},
		{
			in:      Range(0, 0),
			initial: 0,
			want:    0,
		},
		{
			in:      Range(1, 5),
			initial: 10,
			want:    25,
		},
	}

	reducer := func(acc, n int) int {
		return acc + n
	}

	for _, test := range tests {
		got := Reduce(test.in, test.initial, reducer)
		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
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

func iter2AsSlice[T, U any](iter iter.Seq2[T, U]) []pair[T, U] {
	r := []pair[T, U]{}
	for t, u := range iter {
		p := pair[T, U]{first: t, second: u}
		r = append(r, p)
	}
	return r
}

type pair[T, U any] struct {
	first  T
	second U
}
