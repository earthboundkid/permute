package permute

import (
	"iter"
)

// Combinations returns an iterator of indices over the length K combinations of
// an N sized set in lexicograph order. The yielded slice is reused and must be cloned if kept.
//
// E.g. for K = 2 and N = 3, it sets the yielded slice to {0, 1},
// then {0, 2}, and finally {1, 2}.

func Combinations(n, k int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if k > n {
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
