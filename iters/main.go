package main

import "iter"

func main() {}

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
