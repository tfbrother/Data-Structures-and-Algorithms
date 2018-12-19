package graph

// min-heap
// implemented heap.Interface
type EdgeSlice []Edge

func (e EdgeSlice) Len() int { return len(e) }
func (e EdgeSlice) Less(i, j int) bool {
	return e[i].Weight() < e[j].Weight()
}
func (e EdgeSlice) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (n *EdgeSlice) Pop() interface{} {
	ret := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return ret
}

func (n *EdgeSlice) Push(x interface{}) {
	d := x.(Edge)
	*n = append(*n, d)
}
