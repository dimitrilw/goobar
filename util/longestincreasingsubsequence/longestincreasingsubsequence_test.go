package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	t.Run("3 4 7 4", func(t *testing.T) {
		seq := []int{3, 4, 7, 4}
		got := LongestIncreasingSubsequence(seq)
		assert.DeepEqual(t, got, 3)
	})
	t.Run("3 1 2", func(t *testing.T) {
		seq := []int{3, 1, 2}
		got := LongestIncreasingSubsequence(seq)
		assert.DeepEqual(t, got, 2)
	})
	t.Run("1 1 1 1", func(t *testing.T) {
		seq := []int{1, 1, 1, 1}
		got := LongestIncreasingSubsequence(seq)
		assert.DeepEqual(t, got, 1)
	})
	t.Run("3 2 1", func(t *testing.T) {
		seq := []int{3, 2, 1}
		got := LongestIncreasingSubsequence(seq)
		assert.DeepEqual(t, got, 1)
	})
}

// =============================================================================
// Examples

func ExampleLongestIncreasingSubsequence() {
	seq := []int{3, 4, 7, 4}
	fmt.Println(LongestIncreasingSubsequence(seq))
	// Output: 3
}

// =============================================================================
// Benchmarks

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq := []int{3, 4, 7, 4}
		LongestIncreasingSubsequence(seq)
	}
}
