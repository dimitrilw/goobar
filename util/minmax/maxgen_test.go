package minmax

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
	"gotest.tools/v3/assert"
)

func tcRunMaxGen[C constraints.Integer](t *testing.T, tc testCase[C]) {
	t.Helper()
	t.Run(tc.desc, func(t *testing.T) {
		got := MaxGen(tc.give...)
		assert.Equal(t, got, tc.want)
	})
}

func TestMaxGen(t *testing.T) {
	tcRunMaxGen(t, testCase[int]{
		desc: "type int",
		give: []int{1, 2, 3},
		want: int(3),
	})

	tcRunMaxGen(t, testCase[int8]{
		desc: "type int8",
		give: []int8{1, 2, 3},
		want: int8(3),
	})

	tcRunMaxGen(t, testCase[int16]{
		desc: "type int16",
		give: []int16{1, 2, 3},
		want: int16(3),
	})

	tcRunMaxGen(t, testCase[int32]{
		desc: "type int32",
		give: []int32{1, 2, 3},
		want: int32(3),
	})

	tcRunMaxGen(t, testCase[int64]{
		desc: "type int64",
		give: []int64{1, 2, 3},
		want: int64(3),
	})

	tcRunMaxGen(t, testCase[uint]{
		desc: "type uint",
		give: []uint{1, 2, 3},
		want: uint(3),
	})

	tcRunMaxGen(t, testCase[uint8]{
		desc: "type uint8",
		give: []uint8{1, 2, 3},
		want: uint8(3),
	})

	tcRunMaxGen(t, testCase[uint16]{
		desc: "type uint16",
		give: []uint16{1, 2, 3},
		want: uint16(3),
	})

	tcRunMaxGen(t, testCase[uint32]{
		desc: "type uint32",
		give: []uint32{1, 2, 3},
		want: uint32(3),
	})

	tcRunMaxGen(t, testCase[uint64]{
		desc: "type uint64",
		give: []uint64{1, 2, 3},
		want: uint64(3),
	})

	// not sure how to test for type 'uintptr'
	// tcRunMaxGen(t, testCase[uintptr]{
	// 	desc: "type uintptr",
	// 	give: []uintptr{ 1, 2, 3 },
	// 	want: uint64(3),
	// })
}

func TestMaxGen_TypeInferred(t *testing.T) {
	t.Run("type inferred", func(t *testing.T) {
		got := MaxGen(1, 2, 3)
		want := int(3)
		assert.Equal(t, got, want)
	})
}

// =============================================================================
// Examples

func ExampleMaxGen() {
	fmt.Println(MaxGen(1, 2, 3))
	// Output: 3
}

// =============================================================================
// Benchmarks

func BenchmarkMaxGen(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		MaxGen(nums...)
	}
}
