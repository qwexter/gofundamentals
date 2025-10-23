package main

import (
	"iter"
)

func main() {}

// TODO: check the logic in standard lib implementation, how they work with "empty" range.
func Range(start, end int) iter.Seq[int] {
	return func(yeild func(int) bool) {
		for i := start; i < end; i++ {
			if !yeild(i) {
				return
			}
		}
	}
}

func Filter[T any](iter iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yeild func(T) bool) {
		for v := range iter {
			if predicate(v) {
				if !yeild(v) {
					return
				}
			}
		}
	}
}

// Map iterator
// Create an iterator that transforms each value using a mapping function
// Example: Map(Range(1, 3), func(n int) int { return n * 2 })
// should yield: 2, 4, 6
func Map[T, U any](seq iter.Seq[T], mapper func(T) U) iter.Seq[U] {
	return func(yeild func(U) bool) {
		for v := range seq {
			if !yeild(mapper(v)) {
				return
			}
		}
	}
}

// Take - an iterator that yields only the first n elements
// Example: Take(Range(1, 100), 3) should yield: 1, 2, 3
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yeild func(T) bool) {
		if n <= 0 {
			return
		}
		i := 0
		for v := range seq {
			if i >= n || !yeild(v) {
				return
			}
			i++
		}
	}
}

// Cycle Iterator - create an iterator that repeats a slice infinitely
// Example: Cycle([]int{1, 2, 3}) yields: 1, 2, 3, 1, 2, 3, 1, 2, 3, ...
func Cycle[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		if len(items) == 0 {
			return
		}
		for {
			for _, v := range items {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Enumerate - Key-Value Iterator (Seq2)
// Create an iterator that yields both index and value
// Example: Enumerate([]string{"a", "b", "c"}) yields: (0, "a"), (1, "b"), (2, "c")
func Enumerate[T any](items []T) iter.Seq2[int, T] {
	return func(yeild func(int, T) bool) {
		if len(items) <= 0 {
			return
		}
		for i, v := range items {
			if !yeild(i, v) {
				return
			}
		}
	}
}

// Zip Iterator (Seq2) - Create an iterator that combines two sequences into pairs
// Stops when the shorter sequence ends
// Example: Zip(Range(1, 3), Range(10, 15)) yields: (1, 10), (2, 11), (3, 12)
func Zip[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	return func(yeild func(T, U) bool) {
		n1, s1 := iter.Pull(seq1)
		n2, s2 := iter.Pull(seq2)
		defer s1()
		defer s2()
		for {
			v1, ok1 := n1()
			v2, ok2 := n2()
			if !ok1 || !ok2 {
				return
			}
			if !yeild(v1, v2) {
				return
			}
		}
	}
}

// Flatten Iterator
// Create an iterator that flattens a sequence of sequences
// Example: Flatten([][]int{{1, 2}, {3, 4}, {5}}) yields: 1, 2, 3, 4, 5 TODO: rewrite example to work with seq
func Flatten[T any](seqs iter.Seq[[]T]) iter.Seq[T] {
	return func(yeild func(T) bool) {
		for s := range seqs {
			for _, v := range s {
				if !yeild(v) {
					return
				}
			}
		}
	}
}

// Chunk Iterator
// Create an iterator that groups elements into chunks of size n
// Example: Chunk(Range(1, 7), 3) yields: [1,2,3], [4,5,6], [7]
func Chunk[T int](seq iter.Seq[T], size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if size <= 0 {
			return
		}
		chunk := make([]T, 0, size)

		for v := range seq {
			chunk = append(chunk, v)
			if len(chunk) == size {
				if !yield(chunk) {
					return
				}
				chunk = make([]T, 0, size)
			}
		}
		if len(chunk) > 0 {
			if !yield(chunk) {
				return
			}
		}
	}
}

// Reduce/Fold
// Not an iterator, but useful for consuming them
// Reduce a sequence to a single value using an accumulator function
// Example: Reduce(Range(1, 5), 0, func(acc, n int) int { return acc + n })
// should return: 15 (sum of 1+2+3+4+5)
func Reduce[T, U any](seq iter.Seq[T], initial U, reducer func(U, T) U) U {
	r := initial
	for v := range seq {
		r = reducer(r, v)
	}
	return r
}
