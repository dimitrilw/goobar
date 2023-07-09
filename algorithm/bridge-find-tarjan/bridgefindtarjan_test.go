package bridgefindtarjan

// cSpell:ignore costi

import (
	"testing"

	"gotest.tools/v3/assert"
)

/*
LeetCode 1192: Critical Connections in a Network

https://leetcode.com/problems/critical-connections-in-a-network/

There are n servers numbered from 0 to n - 1 connected by undirected server-to-server
connections forming a network where connections[i] = [ai, bi] represents a
connection between servers ai and bi. Any server can reach other servers directly
or indirectly through the network.

A critical connection is a connection that, if removed, will make some servers
unable to reach some other server.

Return all critical connections in the network in any order.
*/

// criticalConnections

func TestCriticalConnections(t *testing.T) {
	/* TEST CASE 1

	Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
	Output: [[1,3]]
	Explanation: [[3,1]] is also accepted.
	*/
	t.Run("test case 1", func(t *testing.T) {
		got := criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}})
		want := [][]int{{1, 3}}
		assert.DeepEqual(t, got, want)
	})
	/* TEST CASE 2

	Input: n = 2, connections = [[0,1]]
	Output: [[0,1]]
	*/
	t.Run("test case 2", func(t *testing.T) {
		got := criticalConnections(2, [][]int{{0, 1}})
		want := [][]int{{0, 1}}
		assert.DeepEqual(t, got, want)
	})
}
