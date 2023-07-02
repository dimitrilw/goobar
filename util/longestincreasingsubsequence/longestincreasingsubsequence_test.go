package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
	"gotest.tools/v3/assert"
)

type testCase[C constraints.Ordered] struct {
	desc string
	give [3]C
}

func TestLISString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		got := LongestIncreasingSubsequenceString("azbycxdwe") // cSpell:disable-line
		assert.Equal(t, got, 5)
	})
}

func tcRun[C constraints.Ordered](t *testing.T, tc testCase[C]) {
	t.Helper()
	t.Run(fmt.Sprintf("%s 0 1 2 1", tc.desc), func(t *testing.T) {
		tArr := []C{tc.give[0], tc.give[1], tc.give[2], tc.give[1]}
		got := LongestIncreasingSubsequence(tArr)
		assert.Equal(t, got, 3)
	})

	t.Run(fmt.Sprintf("%s 2 0 1", tc.desc), func(t *testing.T) {
		tArr := []C{tc.give[2], tc.give[0], tc.give[1]}
		got := LongestIncreasingSubsequence(tArr)
		assert.Equal(t, got, 2)
	})

	t.Run(fmt.Sprintf("%s 1 1 1 1 1", tc.desc), func(t *testing.T) {
		tArr := []C{tc.give[1], tc.give[1], tc.give[1], tc.give[1], tc.give[1]}
		got := LongestIncreasingSubsequence(tArr)
		assert.Equal(t, got, 1)
	})

	t.Run(fmt.Sprintf("%s 2 1 0", tc.desc), func(t *testing.T) {
		tArr := []C{tc.give[2], tc.give[1], tc.give[0]}
		got := LongestIncreasingSubsequence(tArr)
		assert.Equal(t, got, 1)
	})
}

// pipes test cases for various generic types into tcRun"FuncName"
func TestLIS(t *testing.T) {
	tcRun(t, testCase[int]{
		desc: "type int",
		give: [3]int{1, 2, 3},
	})

	tcRun(t, testCase[int64]{
		desc: "type int64",
		give: [3]int64{1, 2, 3},
	})
}

// =============================================================================
// Examples

func ExampleLongestIncreasingSubsequence() {
	seq := []int{1, 2, 3, 2}
	fmt.Println(LongestIncreasingSubsequence(seq))
	// Output: 3
}

func ExampleLongestIncreasingSubsequenceString() {
	fmt.Println(LongestIncreasingSubsequenceString("azbycxdwe")) // cSpell:disable-line

	// Output: 5
}

// =============================================================================
// Benchmarks

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	var seq []int
	for i := 0; i < b.N; i++ {
		seq = append(seq, i)
		LongestIncreasingSubsequence(seq)
	}
}
