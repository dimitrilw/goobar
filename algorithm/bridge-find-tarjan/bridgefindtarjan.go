package bridgefindtarjan

/*  Tarjan's Bridge-Finding Algorithm

    This is a rough implementation of Tarjan's algo, as understood after reading Wikipedia.
    https://en.wikipedia.org/wiki/Bridge_%28graph_theory%29#Tarjan.27s_Bridge-finding_algorithm

    This algo allows a single-pass calculation where each node is visited only one time.
*/

const (
	UNVISITED = -1
)

type TarjanNodeData struct {
	adjList []int // contains other nodes' original node IDs as given
	// in the original function call's args
	tarjanID int // Tarjan ID, which is an incrementing number independent of node ID
	lowTID   int // lowest Tarjan ID this node loops to
}

func criticalConnections(numNodes int, connections [][]int) [][]int {
	nodes := make([]TarjanNodeData, numNodes)
	for i := range nodes {
		nodes[i] = TarjanNodeData{[]int{}, UNVISITED, UNVISITED}
	}
	for _, edge := range connections {
		a, b := edge[0], edge[1]
		nodes[a].adjList = append(nodes[a].adjList, b)
		nodes[b].adjList = append(nodes[b].adjList, a)
	}

	nextTarjanID := 0
	res := [][]int{}

	var dfsTarjan func(int, int)
	dfsTarjan = func(n, from int) {
		nodes[n].tarjanID = nextTarjanID
		nodes[n].lowTID = nextTarjanID
		nextTarjanID++

		// nn = neighbor node
		for _, nn := range nodes[n].adjList {
			// don't retrace the edge that led us to this node
			if nn == from {
				continue
			}

			if nodes[nn].lowTID == UNVISITED {
				dfsTarjan(nn, n)
				nodes[n].lowTID = min(nodes[n].lowTID, nodes[nn].lowTID)
				if nodes[n].tarjanID < nodes[nn].lowTID {
					res = append(res, []int{n, nn})
				}
			} else {
				nodes[n].lowTID = min(nodes[n].lowTID, nodes[nn].tarjanID)
			}
		}
	}
	dfsTarjan(0, 0)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
