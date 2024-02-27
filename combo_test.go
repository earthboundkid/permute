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

func TestCombinationIndices(t *testing.T) {
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
		got := collect(permute.CombinationIndices(tc.N, tc.K))
		be.DeepEqual(t, tc.want, got)
	}
}

func collectstring(seq iter.Seq[[]byte]) []string {
	var s []string
	for v := range seq {
		s = append(s, string(v))
	}
	return s
}

func TestStringCombinations(t *testing.T) {
	for _, tc := range []struct {
		S    string
		K    int
		want []string
	}{
		{},
		{"abc", 2, []string{"ab", "ac", "bc"}},
	} {
		got := collectstring(permute.StringCombinations(tc.S, tc.K))
		be.AllEqual(t, tc.want, got)
	}
}

func TestCombinations(t *testing.T) {
	for _, tc := range []struct {
		S    string
		K    int
		want []string
	}{
		{},
		{"abc", 2, []string{"ab", "ac", "bc"}},
	} {
		b := []byte(tc.S)
		got := collectstring(permute.Combinations(b, tc.K))
		be.AllEqual(t, tc.want, got)
	}
}
