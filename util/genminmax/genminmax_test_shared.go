package genminmax

import (
	"golang.org/x/exp/constraints" // doc at https://pkg.go.dev/golang.org/x/exp/constraints
)

type testCase[C constraints.Integer] struct {
	desc string
	give []C
	want C
}
