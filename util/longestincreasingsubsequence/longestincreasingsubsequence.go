package longestincreasingsubsequence

import (
	"sort"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
)

/*
Longest Increasing Subsequence (LIS) is a classic dynamic programming problem.
The solution below is O(n log n) time and O(n) space.

It is worth remembering that other problems can be solved using LIS as a secondary
step. For example, when given paired data (x, y) where you require that both x and y
are increasing, you can sort the data by x and reverse-sort y, then find the LIS of y.

	Example:

		pairs := []struct{ x, y int}{ {4,5}, {4,6}, {6,7}, {2,3}, {1,1} }
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].x == pairs[j].x {
				return pairs[i].y > pairs[j].y
			}
			return pairs[i].x < pairs[j].x
		}

		// pairs = { {1,1}, {2,3}, {4,6}, {4,5}, {6,7} }
		// now use pairs[1] as the input to LIS

		seq := make([]int, len(pairs))
		for i, p := range pairs {
			seq[i] = p.y
		}

		// seq = [1, 3, 6, 5, 7]
		// now just run LIS on seq

		res := LongestIncreasingSubsequence(seq)
*/
func LongestIncreasingSubsequence[T constraints.Ordered](seq []T) int {
	dp := []T{}
	for _, item := range seq {
		idx := sort.Search(len(dp), func(i int) bool { return dp[i] >= item })
		if idx == len(dp) {
			dp = append(dp, item)
		} else {
			dp[idx] = item
		}
	}
	return len(dp)
}

// helper function to handle strings
func LongestIncreasingSubsequenceString(s string) int {
	return LongestIncreasingSubsequence([]byte(s))
}
