package longestincreasingsubsequence

import (
	"sort"
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
func LongestIncreasingSubsequence(nums []int) int {
	dp := []int{}
	for _, n := range nums {
		idx := sort.SearchInts(dp, n)
		if idx == len(dp) {
			dp = append(dp, n)
		} else {
			dp[idx] = n
		}
	}
	return len(dp)
}
