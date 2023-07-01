package goobar

import (
	"math"
)

const (
	MIN = math.MinInt
)

func Max(N ...int) (r int) {
	r = MIN
	for _, n := range N {
		if n > r {
			r = n
		}
	}
	return
}
