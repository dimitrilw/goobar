package goobar

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestVoid(t *testing.T) {
	t.Run("is a void", func(t *testing.T) {
		var v interface{} = Void{}
		_, got := v.(Void)
		assert.Equal(t, got, true)
	})
}

// =============================================================================
// Examples

func ExampleVoid() {
	var v interface{} = Void{}
	_, ok := v.(Void)
	fmt.Println(ok)
	// Output: true
}

// =============================================================================
// Benchmarks

func BenchmarkVoid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Void{}
	}
}
