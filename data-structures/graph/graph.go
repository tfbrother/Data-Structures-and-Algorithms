package graph

// 图类接口定义
type Graph interface {
	NodeCount() int
	EdgeCount() int
	AddEdge(v int, w int) bool
	DepthFirstTraverse(v int)
	BreadthFirstTraverse(v int)
	ResetStatus()
}
