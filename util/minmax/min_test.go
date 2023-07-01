package goobar

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMin(t *testing.T) {
	t.Run("no arg", func(t *testing.T) {
		got := Min()
		want := math.MaxInt
		assert.Equal(t, got, want)
	})
	t.Run("CSV args", func(t *testing.T) {
		got := Min(1, 2, 3)
		assert.Equal(t, got, 1)
	})
	t.Run("slice expansion args", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Min(nums...)
		assert.Equal(t, got, 1)
	})
}

// =============================================================================
// Examples

func ExampleMin() {
	fmt.Println(Min(1, 2, 3))
	// Output: 1
}

// =============================================================================
// Benchmarks

func BenchmarkMin(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		Min(nums...)
	}
}
