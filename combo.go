// Package permute provides a combination sequence iterator. Requires GOEXPERIMENT=rangefunc.
package permute

import (
	"iter"
)

// CombinationIndices returns an iterator of indices
// over the length K combinations of
// an N sized set
// in lexicograph order.
// The yielded slice is reused and must be cloned if kept.
//
// E.g. for N = 3 and K = 2, it sets the yielded slice to
// {0, 1}, then {0, 2}, and finally {1, 2}.
func CombinationIndices(n, k int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if k > n || k < 1 || n < 0 {
			return
		}
		indices := make([]int, k)
		for i := range indices {
			indices[i] = i
		}

		for {
			if !yield(indices) {
				return
			}

			found := false
			i := 0
			for i = k - 1; i >= 0; i-- {
				if indices[i] != i+n-k {
					found = true
					break
				}
			}
			if !found {
				return
			}

			indices[i]++
			for j := i + 1; j < k; j++ {
				indices[j] = indices[j-1] + 1
			}
		}
	}
}

// StringCombinations takes a string
// and returns an iterator over combinations of sub-string combinations of length K.
// The yielded byte slice is reused and must be cloned or turned into a string if kept.
//
// E.g. For "abc", 2; the yielded slice is set to "ab", "ac", and "bc" succesively.
func StringCombinations(s string, k int) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		buf := make([]byte, k)
		for indices := range CombinationIndices(len(s), k) {
			for i, v := range indices {
				buf[i] = s[v]
			}
			if !yield(buf) {
				return
			}
		}
	}
}

// Combinations takes a slice
// and returns an iterator over combinations of sub-slice combinations of length K.
// The yielded slice is reused and must be cloned if kept.
//
// E.g. For []byte("abc"), 2; the yielded slice is set to "ab", "ac", and "bc" succesively.
func Combinations[E any, S ~[]E](s S, k int) iter.Seq[S] {
	return func(yield func(S) bool) {
		buf := make(S, k)
		for indices := range CombinationIndices(len(s), k) {
			for i, v := range indices {
				buf[i] = s[v]
			}
			if !yield(buf) {
				return
			}
		}
	}
}
