package disjointset

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDisjointSet(t *testing.T) {
	t.Run("NewDisjointSet(0) & Len", func(t *testing.T) {
		d := NewDisjointSet(0)
		assert.Equal(t, d.Len(), 0)
	})
	t.Run("NewDisjointSet(100) & Len", func(t *testing.T) {
		d := NewDisjointSet(100)
		assert.Equal(t, d.Len(), 100)
	})
	t.Run("IsValid", func(t *testing.T) {
		d := NewDisjointSet(10)
		assert.Equal(t, d.IsValid(0), true)
		assert.Equal(t, d.IsValid(5), true)
		assert.Equal(t, d.IsValid(9), true)
		assert.Equal(t, d.IsValid(10), false)
		assert.Equal(t, d.IsValid(99), false)
	})
	t.Run("Add", func(t *testing.T) {
		d := NewDisjointSet(0)
		assert.Equal(t, d.Len(), 0)
		var id int

		id = d.Add()
		assert.Equal(t, id, 0)
		assert.Equal(t, d.Len(), 1)

		id = d.Add()
		assert.Equal(t, id, 1)
		assert.Equal(t, d.Len(), 2)
	})
	t.Run("SetIDs", func(t *testing.T) {
		d := NewDisjointSet(5)
		got := d.SetIDs()
		want := []int{0, 1, 2, 3, 4}
		assert.DeepEqual(t, got, want)
	})
	t.Run("Union", func(t *testing.T) {
		d := NewDisjointSet(5)
		assert.Equal(t, d.Len(), 5)
		var wasMerged bool

		wasMerged = d.Union(0, 1)
		assert.Equal(t, wasMerged, true)
		assert.Equal(t, d.Len(), 4)

		wasMerged = d.Union(0, 1)
		assert.Equal(t, wasMerged, false)
		assert.Equal(t, d.Len(), 4)
	})
	t.Run("Find", func(t *testing.T) {
		d := NewDisjointSet(5)
		d.Union(0, 1)
		assert.Equal(t, d.Len(), 4)
		assert.Equal(t, d.Find(0), d.Find(1))
	})
	t.Run("Find Invalid", func(t *testing.T) {
		d := NewDisjointSet(5)
		assert.Equal(t, d.Find(9), -1)
	})
	t.Run("Size", func(t *testing.T) {
		d := NewDisjointSet(5)
		d.Union(0, 1)
		assert.Equal(t, d.Len(), 4)
		assert.Equal(t, d.Size(0), 2)
	})
	t.Run("Size Invalid", func(t *testing.T) {
		d := NewDisjointSet(5)
		assert.Equal(t, d.Size(9), 0)
	})
}

// =============================================================================
// Examples

func ExampleDisjointSet() {
	d := NewDisjointSet(3)
	d.Union(0, 1) // true
	d.Union(1, 2) // true
	d.Len()       // 1
	d.Add()       // 3
	d.Add()       // 4
	d.Union(3, 4) // true
	d.Union(3, 4) // false
	d.IsValid(2)  // true
	d.IsValid(9)  // false
	d.Len()       // 2
	d.Find(0)     // 1
	d.Size(4)     // 2
	fmt.Println(d.SetIDs())
	// Output: [1 4]
}

// =============================================================================
// Benchmarks

func BenchmarkDisjointSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDisjointSet(3)
		d.Union(0, 1)
		d.Union(1, 2)
		d.Len()
		d.Add()
		d.Add()
		d.Union(3, 4)
		d.Union(3, 4) // duplicate union; will not process
		d.IsValid(2)
		d.IsValid(9)
		d.Find(0)
		d.Size(4)
		d.SetIDs()
	}
}
