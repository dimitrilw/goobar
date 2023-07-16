package flatten2darray

type Vector2D struct {
	ch      chan int
	next    int
	hasNext bool
}

func NewVector2D(vec [][]int) Vector2D {
	v := Vector2D{make(chan int), 0, false}

	// anonymous function contains channel closure & vec processor
	go func() {
		defer close(v.ch)

		for _, row := range vec {
			for _, num := range row {
				// use of channel blocks continuation of go-func
				// until channel is read
				v.ch <- num
			}
		}
	}()

	// init v.next & v.hasNext
	v.Next()

	return v
}

func (v *Vector2D) Next() int {
	res := v.next
	v.next, v.hasNext = <-v.ch
	return res
}

func (v *Vector2D) HasNext() bool { return v.hasNext }
