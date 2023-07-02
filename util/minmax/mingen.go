package minmax

import (
	"github.com/dimitrilw/goobar/imports/constraints" // "golang.org/x/exp/constraints"
)

// note: GenMin will not handle floats. This is by design.
// There are many edge cases with floats, and it is better to use math.Max instead.

func MinGen[T constraints.Integer](args ...T) T {
	var res T
	if len(args) == 0 {
		return res
	}
	res = args[0]
	for _, a := range args {
		if a < res {
			res = a
		}
	}
	return res
}
