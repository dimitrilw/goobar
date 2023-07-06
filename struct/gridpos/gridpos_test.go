package gridpos

import (
	// "fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func gridTestNewTenXTen() Grid {
	return Grid{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{70, 71, 72, 73, 74, 75, 76, 77, 78, 79},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
	}
}

func TestPos(t *testing.T) {
	t.Run("hand-crafted Pos", func(t *testing.T) {
		p := Pos{1, 2}
		assert.Equal(t, p.r, 1)
		assert.Equal(t, p.c, 2)
	})
	t.Run("NewPosFrom slice", func(t *testing.T) {
		in := []int{1, 2}
		p := NewPosFrom(in)
		assert.Equal(t, p.r, 1)
		assert.Equal(t, p.c, 2)
	})
	t.Run("NewPosFrom array", func(t *testing.T) {
		in := [2]int{1, 2}
		p := NewPosFrom(in)
		assert.Equal(t, p.r, 1)
		assert.Equal(t, p.c, 2)
	})
}

func TestGridIsValid(t *testing.T) {
	g := gridTestNewTenXTen()
	t.Run("isValid yes", func(t *testing.T) {
		p := Pos{1, 1}
		got := g.isValid(p)
		assert.Equal(t, got, true)
	})
	t.Run("isValid no, row too low", func(t *testing.T) {
		p := Pos{-1, 1}
		got := g.isValid(p)
		assert.Equal(t, got, false)
	})
	t.Run("isValid no, row too high", func(t *testing.T) {
		p := Pos{11, 1}
		got := g.isValid(p)
		assert.Equal(t, got, false)
	})
	t.Run("isValid no, col too low", func(t *testing.T) {
		p := Pos{1, -1}
		got := g.isValid(p)
		assert.Equal(t, got, false)
	})
	t.Run("isValid no, col too high", func(t *testing.T) {
		p := Pos{1, 11}
		got := g.isValid(p)
		assert.Equal(t, got, false)
	})
}

func TestGridNeighbors(t *testing.T) {
	var g Grid
	g = gridTestNewTenXTen()
	t.Run("Neighbors, middle cell, no diag", func(t *testing.T) {
		p := Pos{5, 5}
		got := g.Neighbors(p, false)
		want := []Pos{{4, 5}, {6, 5}, {5, 4}, {5, 6}} // up, down, left, right
		for i := range got {
			assert.Equal(t, got[i].r, want[i].r)
			assert.Equal(t, got[i].c, want[i].c)
		}
	})
	t.Run("Neighbors, middle cell, with diag", func(t *testing.T) {
		p := Pos{5, 5}
		got := g.Neighbors(p, true)
		// up, down, left, right, up-left, up-right, down-left, down-right
		want := []Pos{{4, 5}, {6, 5}, {5, 4}, {5, 6}, {4, 4}, {4, 6}, {6, 4}, {6, 6}}
		for i := range got {
			assert.Equal(t, got[i].r, want[i].r)
			assert.Equal(t, got[i].c, want[i].c)
		}
	})
	t.Run("Neighbors, corner cell, no diag", func(t *testing.T) {
		p := Pos{0, 0}
		got := g.Neighbors(p, false)
		want := []Pos{{1, 0}, {0, 1}} // down & right
		for i := range got {
			assert.Equal(t, got[i].r, want[i].r)
			assert.Equal(t, got[i].c, want[i].c)
		}
	})
	t.Run("Neighbors, corner cell, with diag", func(t *testing.T) {
		p := Pos{0, 0}
		got := g.Neighbors(p, true)
		want := []Pos{{1, 0}, {0, 1}, {1, 1}} // down, right, & down-right
		for i := range got {
			assert.Equal(t, got[i].r, want[i].r)
			assert.Equal(t, got[i].c, want[i].c)
		}
	})
	g = Grid{{0}}
	t.Run("Neighbors, single cell", func(t *testing.T) {
		p := Pos{0, 0}
		got := g.Neighbors(p, false)
		want := []Pos{} // no neighbors
		for i := range got {
			assert.Equal(t, got[i].r, want[i].r)
			assert.Equal(t, got[i].c, want[i].c)
		}
	})

}

// =============================================================================
// Examples

// yet to add

// =============================================================================
// Benchmarks

func BenchmarkGridPosFromSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := gridTestNewTenXTen()
		demoStartPoints := [][]int{{0, 0}, {4, 2}, {7, 9}, {9, 9}}
		visited := map[Pos]bool{}
		var p Pos
		q := []Pos{}
		for _, sp := range demoStartPoints {
			p = NewPosFrom(sp)
			q = append(q, p)
		}
		// BFS everything, to give the benchmark something to do
		for len(q) > 0 {
			p, q = q[0], q[1:]
			if visited[p] {
				continue
			}
			visited[p] = true
			q = append(q, g.Neighbors(p, false)...)
		}
	}
}

func BenchmarkGridPosFromArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := gridTestNewTenXTen()
		demoStartPoints := [][2]int{{0, 0}, {4, 2}, {7, 9}, {9, 9}}
		visited := map[Pos]bool{}
		var p Pos
		q := []Pos{}
		for _, sp := range demoStartPoints {
			p = NewPosFrom(sp)
			q = append(q, p)
		}
		// BFS everything, to give the benchmark something to do
		for len(q) > 0 {
			p, q = q[0], q[1:]
			if visited[p] {
				continue
			}
			visited[p] = true
			q = append(q, g.Neighbors(p, false)...)
		}
	}
}
