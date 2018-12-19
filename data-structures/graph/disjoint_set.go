package graph

type DisjointSet struct {
	count  int
	parent map[string]string
	rank   map[string]int
}

// O(h), h is the tree depth
func (u *DisjointSet) IsConnected(a string, b string) bool {
	return u.Find(a) == u.Find(b)
}

// Returns the set the element belongs to
func (u *DisjointSet) Find(a string) string {
	for a != u.parent[a] {
		u.parent[a] = u.parent[u.parent[a]] // Path compression(halving)
		a = u.parent[a]
	}
	return a
}

// Union unions two DisjointSet.
func (u *DisjointSet) Union(a string, b string) {
	aParentId, bParentId := u.Find(a), u.Find(b)
	if aParentId != bParentId {
		if u.rank[aParentId] > u.rank[bParentId] {
			u.parent[bParentId] = aParentId // Union
			u.rank[aParentId] += 1
		} else {
			u.parent[aParentId] = bParentId // Union
			u.rank[bParentId] += 1
		}
	}
}

// If the element isn't already somewhere in there, adds it to the master set and its own tiny set.
func (u *DisjointSet) AddElement(e string) {
	u.parent[e] = e
	u.rank[e] = 1
}

func NewDisjointSet(count int) *DisjointSet {
	u := DisjointSet{
		count:  count,
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}

	return &u
}
