package alphatrie

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAlphaTrie(t *testing.T) {
	t.Run("NewAlphaTrie", func(t *testing.T) {
		at := NewAlphaTrie()
		exists := at != nil
		assert.Equal(t, exists, true)
	})
	t.Run("Insert", func(t *testing.T) {
		at := NewAlphaTrie()
		var err error
		err = at.Insert("cat")
		assert.Equal(t, err, nil)
		err = at.Insert("car")
		assert.Equal(t, err, nil)
		err = at.Insert("catch")
		assert.Equal(t, err, nil)
		err = at.Insert("1NV4L1D")
		hasError := err != nil
		assert.Equal(t, hasError, true)
	})
	t.Run("Search", func(t *testing.T) {
		at := NewAlphaTrie()
		at.Insert("cat")
		at.Insert("car")
		at.Insert("catch")
		var got bool
		got = at.Search("car")
		assert.Equal(t, got, true)
		got = at.Search("bunny")
		assert.Equal(t, got, false)
		got = at.Search("1NV4L1D")
		assert.Equal(t, got, false)
	})
	t.Run("StartsWith", func(t *testing.T) {
		at := NewAlphaTrie()
		at.Insert("cat")
		at.Insert("car")
		at.Insert("catch")
		var got bool
		got = at.StartsWith("ca")
		assert.Equal(t, got, true)
		got = at.StartsWith("bun")
		assert.Equal(t, got, false)
		got = at.StartsWith("1NV4L1D")
		assert.Equal(t, got, false)
	})
}

// =============================================================================
// Examples

func ExampleAlphaTrie() {
	at := NewAlphaTrie()
	at.Insert("cat")
	at.Insert("car")
	at.Insert("catch")

	at.StartsWith("bun") // false
	at.Search("1NV4L1D") // false

	fmt.Println(at.StartsWith("ca"))
	// Output: true
}

// =============================================================================
// Benchmarks

func BenchmarkAlphaTrie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		at := NewAlphaTrie()
		at.Insert("cat")
		at.Insert("car")
		at.Insert("catch")
		at.StartsWith("bun")
		at.Search("1NV4L1D")
		at.StartsWith("ca")
	}
}
