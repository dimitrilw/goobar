package heaptype

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
	"gotest.tools/v3/assert"
)

func TestNewHeap(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		h := NewHeap[int](func(a, b int) bool { return a < b })
		heap.Init(h)
		assert.Equal(t, h.Len(), 0)
		heap.Push(h, 2)
		heap.Push(h, 1)
		heap.Push(h, 3)
		assert.Equal(t, h.Len(), 3)
		assert.Equal(t, h.Peek(), 1)
		got := heap.Pop(h)
		assert.Equal(t, got, 1)
		gotSlice := h.Slice()
		sort.Ints(gotSlice)
		want := []int{2, 3}
		assert.DeepEqual(t, gotSlice, want)
	})
}

// =============================================================================
// test different data types

func lessFnForHeapTest[C constraints.Ordered](a, b C) bool { return a < b }

type testCase[C constraints.Ordered] struct {
	desc  string
	start []C
	push  []C
}

func setupRun[C constraints.Ordered](t *testing.T, start []C) (*Heap[C], func(*testing.T)) {
	t.Log("setup before tcRun")

	t.Logf("NewHeapFrom %v", start)
	h := NewHeapFrom[C](start, lessFnForHeapTest[C])
	heap.Init(h)

	return h, func(t *testing.T) {
		t.Log("teardown after tcRun")
	}
}

func tcRun[C constraints.Ordered](t *testing.T, tc testCase[C]) {
	t.Helper()

	t.Run(fmt.Sprintf("%s Len", tc.desc), func(t *testing.T) {
		h, teardown := setupRun(t, tc.start)
		defer teardown(t)

		got := h.Len()
		assert.Equal(t, got, 3)
	})

	t.Run(fmt.Sprintf("%s Peek", tc.desc), func(t *testing.T) {
		h, teardown := setupRun(t, tc.start)
		defer teardown(t)

		want := make([]C, len(tc.start))
		copy(want, tc.start)
		sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })

		got := h.Peek()
		assert.Equal(t, got, want[0])
	})

	t.Run(fmt.Sprintf("%s Push", tc.desc), func(t *testing.T) {
		h, teardown := setupRun(t, tc.start)
		defer teardown(t)

		for _, v := range tc.push {
			heap.Push(h, v)
		}

		got := h.Len()
		assert.Equal(t, got, 6)
	})

	t.Run(fmt.Sprintf("%s Pop", tc.desc), func(t *testing.T) {
		h, teardown := setupRun(t, tc.start)
		defer teardown(t)

		want := make([]C, len(tc.start))
		copy(want, tc.start)
		sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })

		got := heap.Pop(h)
		assert.Equal(t, got, want[0])
		got = heap.Pop(h)
		assert.Equal(t, got, want[1])
		got = heap.Pop(h)
		assert.Equal(t, got, want[2])
	})

	t.Run(fmt.Sprintf("%s Slice", tc.desc), func(t *testing.T) {
		h, teardown := setupRun(t, tc.start)
		defer teardown(t)

		gotSlice := h.Slice()
		sort.Slice(gotSlice, func(i, j int) bool { return gotSlice[i] < gotSlice[j] })
		want := tc.start
		sort.Slice(want, func(i, j int) bool { return gotSlice[i] < gotSlice[j] })
		assert.DeepEqual(t, gotSlice, want)
	})
}

func TestNewHeapFrom_TestCases(t *testing.T) {
	tcRun(t, testCase[int]{
		desc:  "int",
		start: []int{5, 4, 6},
		push:  []int{2, 1, 3},
	})

	tcRun(t, testCase[string]{
		desc:  "string",
		start: []string{"e", "d", "f"},
		push:  []string{"b", "a", "c"},
	})

	tcRun(t, testCase[rune]{
		desc:  "rune",
		start: []rune{'e', 'd', 'f'},
		push:  []rune{'b', 'a', 'c'},
	})
}

// =============================================================================
// Examples

func ExampleHeap() {
	nums := []int{2, 1, 3}
	h := NewHeapFrom(nums, lessFnForHeapTest[int])
	heap.Init(h)
	fmt.Println(heap.Pop(h))
	// Output: 1
}

// =============================================================================
// Benchmarks

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewHeapFrom([]int{1, 2, 3}, lessFnForHeapTest[int])
		s.Len()
		s.Push(5)
		s.Push(4)
		s.Push(6)
		s.Peek()
		s.Pop()
		s.Slice()
	}
}
