package main

import "iter"

func main() {}

func SliceAsIter[T any](s []T) iter.Seq[T] {
	return func(yeild func(T) bool) {
		for _, v := range s {
			if !yeild(v) {
				return
			}
		}
	}
}

func Range(start, end int) iter.Seq[int] {
	return func(yeild func(int) bool) {
		for i := start; i <= end; i++ {
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
	i := 0
	return func(yeild func(T) bool) {
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
