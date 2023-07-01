package genheap

import (
	// "container/heap"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
)

type heapLessFn[T constraints.Ordered] func(T, T) bool

type Heap[T constraints.Ordered] struct {
	lessFn heapLessFn[T]
	s      []T
}

func NewHeap[T constraints.Ordered](lessFn heapLessFn[T]) *Heap[T] {
	return &Heap[T]{lessFn: lessFn}
}
func NewHeapFrom[T constraints.Ordered](s []T, lessFn heapLessFn[T]) *Heap[T] {
	return &Heap[T]{lessFn, s}
}

func (h *Heap[T]) Len() int { return len(h.s) }
func (h *Heap[T]) Less(i, j int) bool {
	return h.lessFn(h.s[i], h.s[j])
}
func (h *Heap[T]) Swap(i, j int) { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *Heap[T]) Push(x any)    { h.s = append(h.s, x.(T)) }
func (h *Heap[T]) Pop() (item any) {
	n := len(h.s)
	item, h.s = h.s[n-1], h.s[:n-1]
	return
}
func (h *Heap[T]) Peek() T    { return h.s[0] }
func (h *Heap[T]) Slice() []T { return h.s }
