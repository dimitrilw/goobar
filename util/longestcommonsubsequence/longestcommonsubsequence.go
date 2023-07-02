package longestcommonsubsequence

import (
	"github.com/dimitrilw/goobar/util/minmax"
)

var Max = minmax.Max // GOOBAR: add Max function

// Longest common subsequence (LCS) generic function for comparing
// two slices of comparable items.
//
// Function uses dp tabulation and is space optimized by working in 1d array
// with vars for other cells-of-interest (from normal 2d dp) & also for readability.
func LongestCommonSubsequence[T comparable](lgSlice, smSlice []T) int {
	if len(lgSlice) < len(smSlice) {
		lgSlice, smSlice = smSlice, lgSlice
	}

	// dp contains a single array of the last "row" of calculations (were this a traditional
	// 2d dp matrix)), which we will cell-by-cell overwrite with the next "row" of calcs
	dp := make([]int, len(smSlice))

	for _, lgItem := range lgSlice {
		// for each new "row" of traditional 2d dp matrix, reset left & diag to 0-val
		left, diag := 0, 0
		for i, smItem := range smSlice {
			// we *could* remove the up variable, but the memory optimization
			// gained is negligible; meanwhile, the readability gain is substantial
			up := dp[i]

			if smItem == lgItem {
				dp[i] = diag + 1
			} else {
				dp[i] = Max(dp[i], left)
			}

			left, diag = dp[i], up
		}
	}
	return dp[len(dp)-1]
}

// helper function to handle strings
func LongestCommonSubsequenceString(lgStr, smStr string) int {
	return LongestCommonSubsequence([]byte(lgStr), []byte(smStr))
}
