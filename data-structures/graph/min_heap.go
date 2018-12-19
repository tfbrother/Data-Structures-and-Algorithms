package graph

// min-heap
// implemented heap.Interface
type EdgeSlice []Edge

func (e EdgeSlice) Len() int { return len(e) }
func (e EdgeSlice) Less(i, j int) bool {
	return e[i].Weight() < e[j].Weight()
}
func (e EdgeSlice) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e *EdgeSlice) Pop() interface{} {
	ret := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return ret
}

func (e *EdgeSlice) Push(x interface{}) {
	d := x.(Edge)
	*e = append(*e, d)
}
