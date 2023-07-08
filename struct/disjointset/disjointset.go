package disjointset

/*
Used for merging sets of data as commonalities are found.

Example use:

	Here is a list of coordinates of interest. Add a new coordinate & if it is touching a neighbor,
	then merge them into one set. If it is touching multiple, merge them all, one-by-one.

	Find loops in a graph by trying to merge nodes & if `DisjointSet.Union()` returns `false`,
	then you know you hit a loop because they are already in the same set.

	Social network groups merging together into a single set.

...and more.

While Len and Size are not typically used, they are provided for convenience.
I encounter the requirement to calculate these values often enough that I decided
to include them. The overhead is small.  Also, if I want a really lean implementation,
it is trivial to prune these features out.
*/

// DisjointSet represents a Disjoint Set (aka Union Find).
// Zero value of DisjointSet is one with 0 elements.
type DisjointSet struct {
	parents, ranks, sizes []int

	numSets int
}

// New returns an initialized disjoint set of given size.
func NewDisjointSet(size int) *DisjointSet {
	d := &DisjointSet{
		parents: make([]int, size),
		ranks:   make([]int, size),
		sizes:   make([]int, size),
		numSets: size,
	}
	for i := 0; i < size; i++ {
		d.parents[i] = i
		d.sizes[i] = 1
	}
	return d
}

// Add adds new element to DisjointSet and returns that element's ID.
func (d *DisjointSet) Add() (id int) {
	id = len(d.parents)
	d.parents = append(d.parents, id)
	d.ranks = append(d.ranks, 0)
	d.sizes = append(d.sizes, 1)
	d.numSets++
	return
}

// Len returns number of sets in DisjointSet.
func (d *DisjointSet) Len() int { return d.numSets }

// IsValid checks if the given ID is a valid int
// for DisjointSet's current state.
func (d *DisjointSet) IsValid(id int) bool {
	return id >= 0 && id < len(d.parents)
}

// SetIDs returns a list of set IDs where each ID is the parent node.
func (d *DisjointSet) SetIDs() []int {
	res := []int{}
	for i, p := range d.parents {
		if i == p {
			res = append(res, i)
		}
	}
	return res
}

// Find returns root element of set containing given element.
func (d *DisjointSet) Find(id int) int {
	if !d.IsValid(id) {
		return -1
	}
	if d.parents[id] != id {
		d.parents[id] = d.Find(d.parents[id])
	}
	return d.parents[id]
}

// Size returns size of set containing given element.
func (d *DisjointSet) Size(id int) int {
	if !d.IsValid(id) {
		return 0
	}

	return d.sizes[d.Find(id)]
}

/*
Union performs the union of two sets containing given elements
and returns true/false on whether or not a union action was completed.
If the two elements are already in the same set, then no union is performed.

The from/dest order is a matter of personal preference. I find that it makes
code line-up better when I use it like this:

	for pos := range listOfNewPositionsOfInterest {
		posID := d.Add()
		if pos.row == 0 {
			d.Union(posID, topGroupID)
		}
		if pos.row == numRows - 1 {
			d.Union(posID, bottomGroupID)
		}
		if pos.col == 0 {
			d.Union(posID, leftGroupID)
		}
		if pos.col == numCols - 1 {
			d.Union(posID, rightGroupID)
		}
	}

Granted, this assumes I ensure the first merge into the "dest" groups (top, bottom, left, right)
was done with that group's ID as the "dest" ID so that it has the highest rank. After that,
order doesn't matter as much. I just find it easier to read when I do it this way.
*/
func (d *DisjointSet) Union(fromID, destID int) bool {
	fromID = d.Find(fromID)
	destID = d.Find(destID)
	// from ID is already in destination ID ...or vice-versa
	if fromID == destID {
		return false
	}

	if d.ranks[fromID] == d.ranks[destID] {
		d.ranks[destID]++
	}

	if d.ranks[destID] < d.ranks[fromID] {
		// fromID is the higher-ranked ID.
		// User had from/dest flipped, and we'll merge anyway.
		fromID, destID = destID, fromID
	}
	d.parents[fromID] = destID
	d.sizes[destID] += d.sizes[fromID]
	d.numSets--
	return true
}
