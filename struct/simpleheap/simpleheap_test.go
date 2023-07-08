package simpleheap

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

func setupRun(t *testing.T) (*Heap, func(*testing.T)) {
	t.Log("setup, building Heap with START_INTS")

	h := &Heap{}
	heap.Init(h)
	for _, n := range START_INTS {
		heap.Push(h, n)
	}

	return h, func(t *testing.T) {
		t.Log("teardown after test")
	}
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
}

// =============================================================================
// Examples

func ExampleHeap() {
	nums := []int{2, 1, 3}
	h := &Heap{}
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
		h := &Heap{}
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
