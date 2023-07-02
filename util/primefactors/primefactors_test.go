package primefactors

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPrimeFactors(t *testing.T) {
	t.Run("108", func(t *testing.T) {
		got := PrimeFactors(108)
		want := []int{2, 2, 3, 3, 3}
		assert.DeepEqual(t, got, want)
	})
	t.Run("12", func(t *testing.T) {
		got := PrimeFactors(12)
		want := []int{2, 2, 3}
		assert.DeepEqual(t, got, want)
	})
	t.Run("1234567890", func(t *testing.T) {
		// cSpell:ignore nums
		got := PrimeFactors(1234567890)
		want := []int{2, 3, 3, 5, 3607, 3803}
		assert.DeepEqual(t, got, want)
	})
}

// =============================================================================
// Examples

func ExamplePrimeFactors() {
	fmt.Println(PrimeFactors(108))
	// Output: [2 2 3 3 3]
}

// =============================================================================
// Benchmarks

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(i)
	}
}
