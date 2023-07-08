package simpleheap

type Heap []int

func (h Heap) Len() int { return len(h) }

// < makes it a MinHeap
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(i any)        { *h = append(*h, i.(int)) }
func (h *Heap) Pop() (res any) {
	i := len(*h) - 1
	res, *h = (*h)[i], (*h)[:i]
	return res
}
func (h *Heap) Peek() int { return (*h)[0] }
