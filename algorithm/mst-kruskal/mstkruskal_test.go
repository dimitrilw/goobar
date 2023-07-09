package mstkruskal

// cSpell:ignore costi

import (
	// "fmt"
	"testing"

	"gotest.tools/v3/assert"
)

/*
LeetCode 1135: Connecting Cities With Minimum Cost

https://leetcode.com/problems/connecting-cities-with-minimum-cost/

There are n cities labeled from 1 to n. You are given the integer n and an array
connections where connections[i] = [xi, yi, costi] indicates that the cost of
connecting city xi and city yi (bidirectional connection) is costi.

Return the minimum cost to connect all the n cities such that there is at least
one path between each pair of cities. If it is impossible to connect all the n
cities, return -1,

The cost is the sum of the connections' costs used.
*/

func TestMinimumCost(t *testing.T) {
	/* TEST CASE 1

	Input: n = 3, connections = [[1,2,5],[1,3,6],[2,3,1]]
	Output: 6
	Explanation: Choosing any 2 edges will connect all cities so we choose the minimum 2.
	*/
	t.Run("test case 1", func(t *testing.T) {
		got := minimumCost(3, [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}})
		want := 6
		assert.Equal(t, got, want)
	})
	/* TEST CASE 2

	Input: n = 4, connections = [[1,2,3],[3,4,4]]
	Output: -1
	Explanation: There is no way to connect all cities even if all edges are used.
	*/
	t.Run("test case 2", func(t *testing.T) {
		got := minimumCost(4, [][]int{{1, 2, 3}, {3, 4, 4}})
		want := -1
		assert.Equal(t, got, want)
	})
}
