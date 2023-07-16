package flatten2darray

import (
	"testing"

	"gotest.tools/v3/assert"
)

/*
LeetCode 251: Flatten 2D Vector

https://leetcode.com/problems/flatten-2d-vector/

Design an iterator to flatten a 2D vector. It should support the Next and HasNext operations.

Implement the Vector2D class:

Vector2D(int[][] vec) initializes the object with the 2D vector vec.

Next() returns the next element from the 2D vector and moves the pointer one step forward.
You may assume that all the calls to next are valid.

HasNext() returns true if there are still some elements in the vector, and false otherwise.
*/

func TestMinimumCost(t *testing.T) {
	/* TEST CASE 1

	Input
	["Vector2D", "next", "next", "next", "hasNext", "hasNext", "next", "hasNext"]
	[[[[1, 2], [3], [4]]], [], [], [], [], [], [], []]
	Output
	[null, 1, 2, 3, true, true, 4, false]
	Explanation
	Vector2D vector2D = new Vector2D([[1, 2], [3], [4]]);
	vector2D.next();    // return 1
	vector2D.next();    // return 2
	vector2D.next();    // return 3
	vector2D.hasNext(); // return True
	vector2D.hasNext(); // return True
	vector2D.next();    // return 4
	vector2D.hasNext(); // return False
	*/
	t.Run("test case 1", func(t *testing.T) {
		// input idx 0
		Vec2DInput := [][]int{{1, 2}, {3}, {4}}
		v := NewVector2D(Vec2DInput)

		// input idx 1
		got := v.Next()
		assert.Equal(t, got, 1)

		// input idx 2
		got = v.Next()
		assert.Equal(t, got, 2)

		// input idx 3
		got = v.Next()
		assert.Equal(t, got, 3)

		// input idx 4
		gotBool := v.HasNext()
		assert.Equal(t, gotBool, true)

		// input idx 5
		gotBool = v.HasNext()
		assert.Equal(t, gotBool, true)

		// input idx 6
		got = v.Next()
		assert.Equal(t, got, 4)

		// input idx 7
		gotBool = v.HasNext()
		assert.Equal(t, gotBool, false)
	})
}
