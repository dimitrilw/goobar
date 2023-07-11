package heapsimple

// import (
// 	"container/heap"
// )

type Heap []int

func (h Heap) Len() int { return len(h) }

// < = MinHeap
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(a any)   { *h = append(*h, a.(int)) }
func (h *Heap) Pop() (res any) {
	i := len(*h) - 1
	res, *h = (*h)[i], (*h)[:i]
	return
}
func (h Heap) Peek() int { return h[0] }
