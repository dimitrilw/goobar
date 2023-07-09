package heaplessfn

// func maxHeapLessFn(a, b int) bool { return a > b }

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

var (
	START_INTS = [3]int{5, 4, 6}
)

func minHeapLessFn(a, b int) bool { return a < b }
func maxHeapLessFn(a, b int) bool { return a > b }

func setupRun(t *testing.T) (*Heap, func(*testing.T)) {
	t.Log("setup, building Heap with START_INTS")

	h := NewHeapFrom(START_INTS[:], minHeapLessFn)
	heap.Init(h)
	// for _, n := range START_INTS {
	// 	heap.Push(h, n)
	// }

	return h, func(t *testing.T) {
		t.Log("teardown after test")
	}
}

func TestNewHeap(t *testing.T) {
	t.Run("NewHeap", func(t *testing.T) {
		h := NewHeap(maxHeapLessFn)
		heap.Init(h)
		heap.Push(h, 1)
		heap.Push(h, 3)
		heap.Push(h, 2)
		got := h.Len()
		assert.Equal(t, got, 3)
		assert.Equal(t, h.Peek(), 3)
	})
	t.Run("NewHeapFrom", func(t *testing.T) {
		h := NewHeapFrom(START_INTS[:], maxHeapLessFn)
		heap.Init(h)
		got := h.Len()
		assert.Equal(t, got, 3)
		assert.Equal(t, h.Peek(), 6)
	})
}

func TestHeap(t *testing.T) {
	t.Run("Len", func(t *testing.T) {
		h, teardown := setupRun(t)
		defer teardown(t)

		got := h.Len()
		assert.Equal(t, got, 3)
	})

	t.Run("Peek", func(t *testing.T) {
		h, teardown := setupRun(t)
		defer teardown(t)

		want := make([]int, len(START_INTS))
		copy(want, START_INTS[:])
		sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })

		got := h.Peek()
		assert.Equal(t, got, want[0])
	})

	t.Run("Push", func(t *testing.T) {
		h, teardown := setupRun(t)
		defer teardown(t)

		for _, v := range [3]int{2, 1, 3} {
			heap.Push(h, v)
		}

		got := h.Len()
		assert.Equal(t, got, 6)
	})

	t.Run("Pop", func(t *testing.T) {
		h, teardown := setupRun(t)
		defer teardown(t)

		want := make([]int, len(START_INTS))
		copy(want, START_INTS[:])
		sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })

		got := heap.Pop(h)
		assert.Equal(t, got, want[0])
		got = heap.Pop(h)
		assert.Equal(t, got, want[1])
		got = heap.Pop(h)
		assert.Equal(t, got, want[2])
	})
	t.Run("Slice", func(t *testing.T) {
		h, teardown := setupRun(t)
		defer teardown(t)

		want := make([]int, len(START_INTS))
		copy(want, START_INTS[:])

		got := h.Slice()

		assert.DeepEqual(t, got, want)
	})
}

// =============================================================================
// Examples

func ExampleHeap() {
	nums := []int{2, 1, 3}
	h := NewHeap(minHeapLessFn)
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}
	fmt.Println(heap.Pop(h))
	// Output: 1
}

// =============================================================================
// Benchmarks

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewHeap(minHeapLessFn)
		heap.Init(h)
		for _, n := range []int{2, 1, 3} {
			heap.Push(h, n)
		}
		h.Len()
		h.Push(5)
		h.Push(4)
		h.Push(6)
		h.Peek()
		h.Pop()
	}
}
