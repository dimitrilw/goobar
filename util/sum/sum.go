package gensum

import (
	"github.com/dimitrilw/goobar/imports/constraints" // "golang.org/x/exp/constraints"
)

// Use of Ordered allows strings too, but I'm okay with that.
func GenSum[T constraints.Ordered](s []T) (r T) {
	for _, v := range s {
		r += v
	}
	return
}
