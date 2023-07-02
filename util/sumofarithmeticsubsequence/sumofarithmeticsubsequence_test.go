package sumofarithmeticsubsequence

// cSpell:ignore sumofarithmeticsubsequence

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSumOfArithmeticSubsequence(t *testing.T) {
	t.Run("same number for args", func(t *testing.T) {
		got := SumOfArithmeticSubsequence(1, 1)
		want := 1
		assert.Equal(t, got, want)

		got = SumOfArithmeticSubsequence(9, 9)
		want = 9
		assert.Equal(t, got, want)
	})

	t.Run("different numbers for args", func(t *testing.T) {
		got := SumOfArithmeticSubsequence(1, 5)
		want := 15
		assert.Equal(t, got, want)

		got = SumOfArithmeticSubsequence(3, 5)
		want = 12
		assert.Equal(t, got, want)
	})
}

// =============================================================================
// Examples

func ExampleSumOfArithmeticSubsequence() {
	fmt.Println(SumOfArithmeticSubsequence(1, 5))
	// Output: 15
}

// =============================================================================
// Benchmarks

func BenchmarkSumOfArithmeticSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfArithmeticSubsequence(1, i)
	}
}
