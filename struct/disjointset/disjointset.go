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

While Len and Size are not typically used, they are provided for convenience. I encounter the requirement
to calculate these values often enough that I decided to include them. The overhead is small.
Also, if I want a really lean implementation, it is trivial to prune these features out.
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

// IsValid checks if the given ID is a valid int
// for DisjointSet's current state.
func (d *DisjointSet) IsValid(id int) bool {
	return id >= 0 && id < len(d.parents)
}

// Len returns number of sets in DisjointSet.
func (d *DisjointSet) Len() int { return d.numSets }

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

// Union performs the union of two sets containing given elements
// and returns true/false on whether or not a union action was completed.
func (d *DisjointSet) Union(smID, lgID int) bool {
	smID = d.Find(smID)
	lgID = d.Find(lgID)
	// small ID is already in large ID ...or vice-versa
	if smID == lgID {
		return false
	}

	if d.ranks[smID] == d.ranks[lgID] {
		d.ranks[lgID]++
	}

	if d.ranks[lgID] > d.ranks[smID] {
		d.parents[smID] = lgID
		d.sizes[lgID] += d.sizes[smID]
	} else {
		// Small ID is the higher-ranked ID.
		// User had small/large flipped; merge anyway.
		d.parents[lgID] = smID
		d.sizes[smID] += d.sizes[lgID]
	}
	d.numSets--
	return true
}
