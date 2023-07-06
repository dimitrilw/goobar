package gridpos

// Many code challenges present a grid and ask for some sort of navigation
// across the grid. This package provides a simple struct to represent
// that grid and the navigation method I find myself using most often.

type PosInput interface {
	[]int | [2]int
}

type Pos struct{ r, c int }

func NewPosFrom[T PosInput](p T) Pos {
	return Pos{p[0], p[1]}
}

type Grid [][]int // GOOBAR: change as required by the problem; e.g., [][]byte

func (g Grid) isValid(p Pos) bool {
	if p.r < 0 || p.r >= len(g) {
		return false
	}
	if p.c < 0 || p.c >= len(g[p.r]) {
		return false
	}
	return true
}

func (g Grid) Neighbors(p Pos, includeDiag bool) []Pos {
	res := []Pos{}

	dirs := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	for _, d := range dirs {
		np := Pos{p.r + d.r, p.c + d.c}
		if g.isValid(np) {
			res = append(res, np)
		}
	}

	if includeDiag {
		diags := []Pos{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} // UL, UR, DL, DR
		for _, d := range diags {
			np := Pos{p.r + d.r, p.c + d.c}
			if g.isValid(np) {
				res = append(res, np)
			}
		}
	}

	return res
}
