package gensum

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
	"gotest.tools/v3/assert"
)

type testCase[C constraints.Ordered] struct {
	desc string
	give []C
	want C
}

func tcRun[C constraints.Ordered](t *testing.T, tc testCase[C]) {
	t.Helper()
	t.Run(tc.desc, func(t *testing.T) {
		got := GenSum(tc.give)
		assert.Equal(t, got, tc.want)
	})
}

func TestGenSum(t *testing.T) {
	tcRun(t, testCase[int]{
		desc: "type int",
		give: []int{1, 2, 3},
		want: int(6),
	})

	tcRun(t, testCase[int8]{
		desc: "type int8",
		give: []int8{1, 2, 3},
		want: int8(6),
	})

	tcRun(t, testCase[int16]{
		desc: "type int16",
		give: []int16{1, 2, 3},
		want: int16(6),
	})

	tcRun(t, testCase[int32]{
		desc: "type int32",
		give: []int32{1, 2, 3},
		want: int32(6),
	})

	tcRun(t, testCase[int64]{
		desc: "type int64",
		give: []int64{1, 2, 3},
		want: int64(6),
	})

	tcRun(t, testCase[uint]{
		desc: "type uint",
		give: []uint{1, 2, 3},
		want: uint(6),
	})

	tcRun(t, testCase[uint8]{
		desc: "type uint8",
		give: []uint8{1, 2, 3},
		want: uint8(6),
	})

	tcRun(t, testCase[uint16]{
		desc: "type uint16",
		give: []uint16{1, 2, 3},
		want: uint16(6),
	})

	tcRun(t, testCase[uint32]{
		desc: "type uint32",
		give: []uint32{1, 2, 3},
		want: uint32(6),
	})

	tcRun(t, testCase[uint64]{
		desc: "type uint64",
		give: []uint64{1, 2, 3},
		want: uint64(6),
	})

	// not sure how to test for type 'uintptr'
	// tcRun(t, testCase[uintptr]{
	// 	desc: "type uintptr",
	// 	give: []uintptr{ 1, 2, 3 },
	// 	want: uint64(3),
	// })

	tcRun(t, testCase[string]{
		desc: "type string",
		give: []string{"one", "two", "three"},
		want: string("onetwothree"),
	})
}

// =============================================================================
// Examples

func ExampleGenSum() {
	fmt.Println(GenSum([]int{1, 2, 3}))
	// Output: 6
}

// =============================================================================
// Benchmarks

func BenchmarkGenSum(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		GenSum(nums)
	}
}
