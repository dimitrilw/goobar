package heapgeneric

import (
	// "container/heap"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
)

type heapLessFn[T constraints.Ordered] func(T, T) bool

type Heap[T constraints.Ordered] struct {
	lessFn heapLessFn[T]
	data   []T
}

func NewHeap[T constraints.Ordered](lessFn heapLessFn[T]) *Heap[T] {
	return &Heap[T]{lessFn: lessFn}
}
func NewHeapFrom[T constraints.Ordered](data []T, lessFn heapLessFn[T]) *Heap[T] {
	return &Heap[T]{lessFn, data}
}

func (h *Heap[T]) Len() int { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool {
	return h.lessFn(h.data[i], h.data[j])
}
func (h *Heap[T]) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap[T]) Push(a any)    { h.data = append(h.data, a.(T)) }
func (h *Heap[T]) Pop() (res any) {
	i := len(h.data) - 1
	res, h.data = h.data[i], h.data[:i]
	return
}
func (h *Heap[T]) Peek() T    { return h.data[0] }
func (h *Heap[T]) Slice() []T { return h.data }
