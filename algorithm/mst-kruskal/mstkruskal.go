package mstkruskal

import (
	"sort"

	"github.com/dimitrilw/goobar/struct/disjointset"
)

// note: in challenges, some features (e.g., size) can be pruned from DisjointSet

// DisjointSet as part of Kruskal's algo for Minimum Spanning Tree (MST)
func minimumCost(n int, connections [][]int) int {
	d := disjointset.NewDisjointSet(n)

	// sort connections: cheapest first
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	// add connections one-by-one until all nodes are connected
	res := 0
	for _, c := range connections {
		// all done
		if d.Len() < 2 {
			break
		}

		// -1 because the data is 1-indexed
		id0, id1 := d.Find(c[0]-1), d.Find(c[1]-1)

		// skip if nodes are already in same group
		if id0 == id1 {
			continue
		}

		d.Union(id0, id1)
		res += c[2]
	}

	if d.Len() == 1 {
		return res
	}
	return -1
}
