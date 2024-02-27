package permute_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/earthboundkid/permute"
)

func collect[T any, S []T](seq iter.Seq[S]) []S {
	var s []S
	for v := range seq {
		s = append(s, slices.Clone(v))
	}
	return s
}

func TestCombinations(t *testing.T) {
	for _, tc := range []struct {
		N, K int
		want [][]int
	}{
		{0, 0, nil},
		{3, 2, [][]int{
			{0, 1},
			{0, 2},
			{1, 2},
		}},
		{5, 3, [][]int{
			{0, 1, 2},
			{0, 1, 3},
			{0, 1, 4},
			{0, 2, 3},
			{0, 2, 4},
			{0, 3, 4},
			{1, 2, 3},
			{1, 2, 4},
			{1, 3, 4},
			{2, 3, 4},
		}},
	} {
		got := collect(permute.Combinations(tc.N, tc.K))
		be.DeepEqual(t, tc.want, got)
	}
}
