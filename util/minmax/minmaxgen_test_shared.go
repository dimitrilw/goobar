package minmax

import (
	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
)

type testCase[C constraints.Integer] struct {
	desc string
	give []C
	want C
}
