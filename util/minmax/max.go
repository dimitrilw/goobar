package minmax

import (
	"math"
)

const (
	MIN = math.MinInt
)

// note: do not use custom Max function with floats; use math.Max instead
// because there are many special cases with floats

func Max(N ...int) (r int) {
	r = MIN
	for _, n := range N {
		if n > r {
			r = n
		}
	}
	return
}
