package goobar

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMax(t *testing.T) {
	t.Run("no arg", func(t *testing.T) {
		got := Max()
		want := math.MinInt
		assert.Equal(t, got, want)
	})
	t.Run("CSV args", func(t *testing.T) {
		got := Max(1, 2, 3)
		assert.Equal(t, got, 3)
	})
	t.Run("slice expansion args", func(t *testing.T) {
		// cSpell:ignore nums
		nums := []int{1, 2, 3}
		got := Max(nums...)
		assert.Equal(t, got, 3)
	})
}

// =============================================================================
// Examples

func ExampleMax() {
	fmt.Println(Max(1, 2, 3))
	// Output: 3
}

// =============================================================================
// Benchmarks

func BenchmarkMax(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		Max(nums...)
	}
}
