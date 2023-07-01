package set

import (
	"fmt"
	"testing"

	// "github.com/dimitrilw/goobar/imports/constraints"
	"gotest.tools/v3/assert"
)

type posStructForSetTest struct{ r, c int }

type testCase[C comparable] struct {
	desc  string
	start []C
	add   []C
}

func setupSubTest[C comparable](t *testing.T, start []C) (Set[C], func(*testing.T)) {
	t.Log("setup sub test")
	s := NewSet[C](start)
	return s, func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func tcRunSet[C comparable](t *testing.T, tc testCase[C]) {
	t.Helper()

	t.Run(fmt.Sprintf("%s Len", tc.desc), func(t *testing.T) {
		s, teardown := setupSubTest(t, tc.start)
		defer teardown(t)

		got := s.Len()
		assert.Equal(t, got, 3)
	})

	t.Run(fmt.Sprintf("%s Has", tc.desc), func(t *testing.T) {
		s, teardown := setupSubTest(t, tc.start)
		defer teardown(t)

		got := s.Has(tc.start[0])
		assert.Equal(t, got, true)
	})

	t.Run(fmt.Sprintf("%s Add", tc.desc), func(t *testing.T) {
		s, teardown := setupSubTest(t, tc.start)
		defer teardown(t)

		for _, a := range tc.add {
			s.Add(a)
		}
		got := s.Has(tc.add[0])
		assert.Equal(t, got, true)
	})

	t.Run(fmt.Sprintf("%s Del", tc.desc), func(t *testing.T) {
		s, teardown := setupSubTest(t, tc.start)
		defer teardown(t)

		s.Del(tc.start[0])
		got := s.Has(tc.start[0])
		assert.Equal(t, got, false)
	})

	t.Run(fmt.Sprintf("%s Slice", tc.desc), func(t *testing.T) {
		s, teardown := setupSubTest(t, tc.start)
		defer teardown(t)

		got := s.Slice()
		for _, val := range got {
			assert.Equal(t, s.Has(val), true)
		}
	})
}

func TestSet_TestCases(t *testing.T) {
	tcRunSet(t, testCase[int]{
		desc:  "int",
		start: []int{1, 2, 3},
		add:   []int{4, 5, 6},
	})

	tcRunSet(t, testCase[string]{
		desc:  "string",
		start: []string{"a", "b", "c"},
		add:   []string{"d", "e", "f"},
	})

	tcRunSet(t, testCase[rune]{
		desc:  "rune",
		start: []rune{'a', 'b', 'c'},
		add:   []rune{'d', 'e', 'f'},
	})

	tcRunSet(t, testCase[posStructForSetTest]{
		desc:  "posStruct",
		start: []posStructForSetTest{{1, 1}, {2, 2}, {3, 3}},
		add:   []posStructForSetTest{{4, 4}, {5, 5}, {6, 6}},
	})
}

// =============================================================================
// Examples

func ExampleSet() {
	nums := []int{1}
	fmt.Println(NewSet(nums))
	// Output: {map[1:{}]}
}

// =============================================================================
// Benchmarks

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSet([]int{1, 2, 3})
		s.Len()
		s.Add(i)
		s.Has(i)
		s.Del(i)
		s.Add(i)
	}
}
