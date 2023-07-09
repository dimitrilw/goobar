package heapsimple

type Heap []int

func (h Heap) Len() int { return len(h) }

// < = MinHeap
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)   { *h = append(*h, x.(int)) }
func (h *Heap) Pop() (item any) {
	i := len(*h) - 1
	item, *h = (*h)[i], (*h)[:i]
	return
}
func (h Heap) Peek() int { return h[0] }
