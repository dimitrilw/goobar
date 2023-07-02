package minmax

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints" // "github.com/dimitrilw/goobar/imports/constraints"
	"gotest.tools/v3/assert"
)

func tcRunMinGen[C constraints.Integer](t *testing.T, tc testCase[C]) {
	t.Helper()
	t.Run(tc.desc, func(t *testing.T) {
		got := MinGen(tc.give...)
		assert.Equal(t, got, tc.want)
	})
}

func TestMinGen(t *testing.T) {
	tcRunMinGen(t, testCase[int]{
		desc: "type int",
		give: []int{1, 2, 3},
		want: int(1),
	})

	tcRunMinGen(t, testCase[int8]{
		desc: "type int8",
		give: []int8{1, 2, 3},
		want: int8(1),
	})

	tcRunMinGen(t, testCase[int16]{
		desc: "type int16",
		give: []int16{1, 2, 3},
		want: int16(1),
	})

	tcRunMinGen(t, testCase[int32]{
		desc: "type int32",
		give: []int32{1, 2, 3},
		want: int32(1),
	})

	tcRunMinGen(t, testCase[int64]{
		desc: "type int64",
		give: []int64{1, 2, 3},
		want: int64(1),
	})

	tcRunMinGen(t, testCase[uint]{
		desc: "type uint",
		give: []uint{1, 2, 3},
		want: uint(1),
	})

	tcRunMinGen(t, testCase[uint8]{
		desc: "type uint8",
		give: []uint8{1, 2, 3},
		want: uint8(1),
	})

	tcRunMinGen(t, testCase[uint16]{
		desc: "type uint16",
		give: []uint16{1, 2, 3},
		want: uint16(1),
	})

	tcRunMinGen(t, testCase[uint32]{
		desc: "type uint32",
		give: []uint32{1, 2, 3},
		want: uint32(1),
	})

	tcRunMinGen(t, testCase[uint64]{
		desc: "type uint64",
		give: []uint64{1, 2, 3},
		want: uint64(1),
	})

	// not sure how to test for type 'uintptr'
	// tcRunMinGen(t, testCase[uintptr]{
	// 	desc: "type uintptr",
	// 	give: []uintptr{ 1, 2, 3 },
	// 	want: uint64(1),
	// })
}

func TestMinGen_TypeInferred(t *testing.T) {
	t.Run("type inferred", func(t *testing.T) {
		got := MinGen(1, 2, 3)
		want := int(1)
		assert.Equal(t, got, want)
	})
}

// =============================================================================
// Examples

func ExampleMinGen() {
	fmt.Println(MinGen(1, 2, 3))
	// Output: 1
}

// =============================================================================
// Benchmarks

func BenchmarkMinGen(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		MinGen(nums...)
	}
}
