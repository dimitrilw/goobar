package goobar

import (
	"math"
)

const (
	MAX = math.MaxInt
)

func Min(N ...int) (r int) {
	r = MAX
	for _, n := range N {
		if n < r {
			r = n
		}
	}
	return
}
