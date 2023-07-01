package minmax

import (
	"math"
)

const (
	MAX = math.MaxInt
)

// note: do not use custom Min function with floats; use math.Max instead
// because there are many special cases with floats

func Min(N ...int) (r int) {
	r = MAX
	for _, n := range N {
		if n < r {
			r = n
		}
	}
	return
}
