package set

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewSet(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, s.Len(), 0)
		assert.Equal(t, s.Has(1), false)
		s.Add(1)
		assert.Equal(t, s.Has(1), true)
	})
}

// =============================================================================
// test different data types

type posStructForSetTest struct{ r, c int }

type testCase[C comparable] struct {
	desc  string
	start []C
	add   []C
}

func setupRun[C comparable](t *testing.T, start []C) (Set[C], func(*testing.T)) {
	t.Log("setup before tcRun")

	t.Logf("NewSetFrom %v", start)
	s := NewSetFrom[C](start)

	return s, func(t *testing.T) {
		t.Log("teardown after tcRun")
	}
}

func tcRun[C comparable](t *testing.T, tc testCase[C]) {
	t.Helper()

	t.Run(fmt.Sprintf("%s Len", tc.desc), func(t *testing.T) {
		s, teardown := setupRun(t, tc.start)
		defer teardown(t)

		got := s.Len()
		assert.Equal(t, got, 3)
	})

	t.Run(fmt.Sprintf("%s Has", tc.desc), func(t *testing.T) {
		s, teardown := setupRun(t, tc.start)
		defer teardown(t)

		got := s.Has(tc.start[0])
		assert.Equal(t, got, true)
	})

	t.Run(fmt.Sprintf("%s Add", tc.desc), func(t *testing.T) {
		s, teardown := setupRun(t, tc.start)
		defer teardown(t)

		for _, a := range tc.add {
			s.Add(a)
		}
		got := s.Has(tc.add[0])
		assert.Equal(t, got, true)
	})

	t.Run(fmt.Sprintf("%s Del", tc.desc), func(t *testing.T) {
		s, teardown := setupRun(t, tc.start)
		defer teardown(t)

		s.Del(tc.start[0])
		got := s.Has(tc.start[0])
		assert.Equal(t, got, false)
	})

	t.Run(fmt.Sprintf("%s Slice", tc.desc), func(t *testing.T) {
		s, teardown := setupRun(t, tc.start)
		defer teardown(t)

		got := s.Slice()
		for _, val := range got {
			assert.Equal(t, s.Has(val), true)
		}
	})
}

func TestNewSetFrom(t *testing.T) {
	tcRun(t, testCase[int]{
		desc:  "int",
		start: []int{1, 2, 3},
		add:   []int{4, 5, 6},
	})

	tcRun(t, testCase[string]{
		desc:  "string",
		start: []string{"a", "b", "c"},
		add:   []string{"d", "e", "f"},
	})

	tcRun(t, testCase[rune]{
		desc:  "rune",
		start: []rune{'a', 'b', 'c'},
		add:   []rune{'d', 'e', 'f'},
	})

	tcRun(t, testCase[posStructForSetTest]{
		desc:  "posStruct",
		start: []posStructForSetTest{{1, 1}, {2, 2}, {3, 3}},
		add:   []posStructForSetTest{{4, 4}, {5, 5}, {6, 6}},
	})
}

// =============================================================================
// Examples

func ExampleSet() {
	nums := []int{1}
	fmt.Println(NewSetFrom(nums))
	// Output: {map[1:{}]}
}

// =============================================================================
// Benchmarks

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSetFrom([]int{1, 2, 3})
		s.Len()
		s.Add(i)
		s.Has(i)
		s.Del(i)
		s.Add(i)
	}
}
