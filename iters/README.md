package main

import (
	"fmt"
	"iter"
)

// EXERCISE 1: Basic Iterator
// Create an iterator that yields numbers from start to end (inclusive)
// Example: Range(1, 5) should yield: 1, 2, 3, 4, 5
func Range(start, end int) iter.Seq[int] {
	// TODO: Implement this
	return nil
}

// EXERCISE 2: Filter Iterator
// Create an iterator that filters values based on a predicate function
// Example: Filter(Range(1, 10), func(n int) bool { return n%2 == 0 })
// should yield only even numbers: 2, 4, 6, 8, 10
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	// TODO: Implement this
	return nil
}

// EXERCISE 3: Map Iterator
// Create an iterator that transforms each value using a mapping function
// Example: Map(Range(1, 3), func(n int) int { return n * 2 })
// should yield: 2, 4, 6
func Map[T, U any](seq iter.Seq[T], mapper func(T) U) iter.Seq[U] {
	// TODO: Implement this
	return nil
}

// EXERCISE 4: Take Iterator
// Create an iterator that yields only the first n elements
// Example: Take(Range(1, 100), 3) should yield: 1, 2, 3
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	// TODO: Implement this
	return nil
}

// EXERCISE 5: Cycle Iterator
// Create an iterator that repeats a slice infinitely
// Example: Cycle([]int{1, 2, 3}) yields: 1, 2, 3, 1, 2, 3, 1, 2, 3, ...
// Hint: Use this with Take() to avoid infinite loops when testing
func Cycle[T any](items []T) iter.Seq[T] {
	// TODO: Implement this
	return nil
}

// EXERCISE 6: Key-Value Iterator (Seq2)
// Create an iterator that yields both index and value
// Example: Enumerate([]string{"a", "b", "c"}) yields: (0, "a"), (1, "b"), (2, "c")
func Enumerate[T any](items []T) iter.Seq2[int, T] {
	// TODO: Implement this
	return nil
}

// EXERCISE 7: Zip Iterator (Seq2)
// Create an iterator that combines two sequences into pairs
// Stops when the shorter sequence ends
// Example: Zip(Range(1, 3), Range(10, 15)) yields: (1, 10), (2, 11), (3, 12)
func Zip[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	// TODO: Implement this
	return nil
}

// EXERCISE 8: Flatten Iterator
// Create an iterator that flattens a sequence of sequences
// Example: Flatten([][]int{{1, 2}, {3, 4}, {5}}) yields: 1, 2, 3, 4, 5
func Flatten[T any](seqs iter.Seq[[]T]) iter.Seq[T] {
	// TODO: Implement this
	return nil
}

// EXERCISE 9: Chunk Iterator
// Create an iterator that groups elements into chunks of size n
// Example: Chunk(Range(1, 7), 3) yields: [1,2,3], [4,5,6], [7]
func Chunk[T any](seq iter.Seq[T], size int) iter.Seq[[]T] {
	// TODO: Implement this
	return nil
}

// EXERCISE 10: Reduce/Fold
// Not an iterator, but useful for consuming them
// Reduce a sequence to a single value using an accumulator function
// Example: Reduce(Range(1, 5), 0, func(acc, n int) int { return acc + n })
// should return: 15 (sum of 1+2+3+4+5)
func Reduce[T, U any](seq iter.Seq[T], initial U, reducer func(U, T) U) U {
	// TODO: Implement this
	return initial
}

// ==================== SOLUTIONS BELOW ====================
// Try to solve the exercises above before looking at these!
// =========================================================

// SOLUTION 1
func RangeSolution(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// SOLUTION 2
func FilterSolution[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// SOLUTION 3
func MapSolution[T, U any](seq iter.Seq[T], mapper func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

// SOLUTION 4
func TakeSolution[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for v := range seq {
			if count >= n {
				return
			}
			if !yield(v) {
				return
			}
			count++
		}
	}
}

// SOLUTION 5
func CycleSolution[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		if len(items) == 0 {
			return
		}
		for {
			for _, item := range items {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// SOLUTION 6
func EnumerateSolution[T any](items []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range items {
			if !yield(i, v) {
				return
			}
		}
	}
}

// SOLUTION 7
func ZipSolution[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		
		for v1 := range seq1 {
			v2, ok := next2()
			if !ok {
				return
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// SOLUTION 8
func FlattenSolution[T any](seqs iter.Seq[[]T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for slice := range seqs {
			for _, item := range slice {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// SOLUTION 9
func ChunkSolution[T any](seq iter.Seq[T], size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
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

// SOLUTION 10
func ReduceSolution[T, U any](seq iter.Seq[T], initial U, reducer func(U, T) U) U {
	acc := initial
	for v := range seq {
		acc = reducer(acc, v)
	}
	return acc
}

// ==================== TEST EXAMPLES ====================

func main() {
	fmt.Println("Go Iterator Exercises")
	fmt.Println("=====================\n")

	// Test your implementations here!
	// Uncomment to test (replace with your function names)

	// Example Test 1: Range
	// fmt.Println("Range(1, 5):")
	// for v := range RangeSolution(1, 5) {
	// 	fmt.Println(v)
	// }

	// Example Test 2: Filter
	// fmt.Println("\nFilter even numbers from Range(1, 10):")
	// for v := range FilterSolution(RangeSolution(1, 10), func(n int) bool { return n%2 == 0 }) {
	// 	fmt.Println(v)
	// }

	// Example Test 3: Map
	// fmt.Println("\nMap (double) Range(1, 5):")
	// for v := range MapSolution(RangeSolution(1, 5), func(n int) int { return n * 2 }) {
	// 	fmt.Println(v)
	// }

	// Example Test 4: Take
	// fmt.Println("\nTake(Range(1, 100), 3):")
	// for v := range TakeSolution(RangeSolution(1, 100), 3) {
	// 	fmt.Println(v)
	// }

	// Example Test 5: Cycle
	// fmt.Println("\nTake(Cycle([1, 2, 3]), 7):")
	// for v := range TakeSolution(CycleSolution([]int{1, 2, 3}), 7) {
	// 	fmt.Println(v)
	// }

	// Example Test 6: Enumerate
	// fmt.Println("\nEnumerate([\"a\", \"b\", \"c\"]):")
	// for i, v := range EnumerateSolution([]string{"a", "b", "c"}) {
	// 	fmt.Printf("(%d, %s)\n", i, v)
	// }

	// Example Test 7: Zip
	// fmt.Println("\nZip(Range(1, 3), Range(10, 15)):")
	// for v1, v2 := range ZipSolution(RangeSolution(1, 3), RangeSolution(10, 15)) {
	// 	fmt.Printf("(%d, %d)\n", v1, v2)
	// }

	// Example Test 8: Flatten
	// fmt.Println("\nFlatten([[1, 2], [3, 4], [5]]):")
	// slices := func(yield func([]int) bool) {
	// 	yield([]int{1, 2})
	// 	yield([]int{3, 4})
	// 	yield([]int{5})
	// }
	// for v := range FlattenSolution(slices) {
	// 	fmt.Println(v)
	// }

	// Example Test 9: Chunk
	// fmt.Println("\nChunk(Range(1, 7), 3):")
	// for chunk := range ChunkSolution(RangeSolution(1, 7), 3) {
	// 	fmt.Println(chunk)
	// }

	// Example Test 10: Reduce
	// fmt.Println("\nReduce(Range(1, 5), 0, sum):")
	// sum := ReduceSolution(RangeSolution(1, 5), 0, func(acc, n int) int { return acc + n })
	// fmt.Println(sum)
}
