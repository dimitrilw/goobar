package longestcommonsubsequencegen

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

type testCase[C comparable] struct {
	desc   string
	giveLg []C
	giveSm []C
	want   int
}

type testLCS_Pos struct{ r, c int }

func TestLCSString(t *testing.T) {
	lgStr, smStr := "abcde", "ace"
	t.Run("string run clean", func(t *testing.T) {
		got := LongestCommonSubsequenceString(lgStr, smStr)
		assert.Equal(t, got, 3)
	})

	t.Run("string run flipped", func(t *testing.T) {
		got := LongestCommonSubsequenceString(smStr, lgStr)
		assert.Equal(t, got, 3)
	})
}

func tcRunLCSGen[C comparable](t *testing.T, tc testCase[C]) {
	t.Helper()
	t.Run(fmt.Sprintf("%s run clean", tc.desc), func(t *testing.T) {
		got := LongestCommonSubsequenceGen(tc.giveLg, tc.giveSm)
		assert.Equal(t, got, tc.want)
	})

	t.Run(fmt.Sprintf("%s run flipped", tc.desc), func(t *testing.T) {
		got := LongestCommonSubsequenceGen(tc.giveSm, tc.giveLg)
		assert.Equal(t, got, tc.want)
	})
}

func TestLCSGen_TestCases(t *testing.T) {
	tcRunLCSGen(t, testCase[int]{
		desc:   "type int",
		giveLg: []int{1, 2, 3, 4, 5},
		giveSm: []int{1, 3, 5},
		want:   3,
	})

	tcRunLCSGen(t, testCase[int64]{
		desc:   "type int64",
		giveLg: []int64{1, 2, 3, 4, 5},
		giveSm: []int64{1, 3, 5},
		want:   3,
	})

	tcRunLCSGen(t, testCase[testLCS_Pos]{
		desc:   "type Pos struct",
		giveLg: []testLCS_Pos{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
		giveSm: []testLCS_Pos{{1, 1}, {3, 3}, {5, 5}},
		want:   3,
	})

	// nobody would run this on bools, but I'm also playing with test cases here
	tcRunLCSGen(t, testCase[bool]{
		desc:   "type bool",
		giveLg: []bool{true, false, true, false, true},
		giveSm: []bool{true, true, true},
		want:   3,
	})
}

// =============================================================================
// Examples

func ExampleLongestCommonSubsequenceGen() {
	lg := []int{1, 2, 3, 4, 5}
	sm := []int{1, 3, 5}
	fmt.Println(LongestCommonSubsequenceGen(lg, sm))
	// Output: 3
}

func ExampleLongestCommonSubsequenceString() {
	lg := "abcde"
	sm := "ace"
	fmt.Println(LongestCommonSubsequenceString(lg, sm))
	// Output: 3
}

// =============================================================================
// Benchmarks

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	var lg, sm []int
	for i := 0; i < b.N; i++ {
		lg = append(lg, i)
		if i%2 == 0 {
			sm = append(sm, i)
		}
		LongestCommonSubsequenceGen(lg, sm)
	}
}
