package heaplessfn

// import (
// 	"container/heap"
// )

type heapLessFn func(int, int) bool

// func minHeapLessFn(a, b int) bool { return a < b }
// func maxHeapLessFn(a, b int) bool { return a > b }

// Heap type allows for quick implementation of both min/max heaps
type Heap struct {
	lessFn heapLessFn
	data   []int
}

func NewHeap(lessFn heapLessFn) *Heap {
	return &Heap{lessFn: lessFn}
}
func NewHeapFrom(data []int, lessFn heapLessFn) *Heap {
	return &Heap{lessFn, data}
}

func (h *Heap) Len() int { return len(h.data) }
func (h *Heap) Less(i, j int) bool {
	return h.lessFn(h.data[i], h.data[j])
}
func (h *Heap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap) Push(a any)    { h.data = append(h.data, a.(int)) }
func (h *Heap) Pop() (res any) {
	i := len(h.data) - 1
	res, h.data = h.data[i], h.data[:i]
	return
}
func (h *Heap) Peek() int    { return h.data[0] }
func (h *Heap) Slice() []int { return h.data }
