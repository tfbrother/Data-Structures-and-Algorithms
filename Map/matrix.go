package Map

import (
	"fmt"
)

// 图结点类型
type Node struct {
	value    string // 顶点的值
	isAccess bool   // 顶点是否被访问
}

func NewNode(value string) Node {
	return Node{value: value}
}

// 边(用于最小生成数使用)
type Edge struct {
	nodeIndexA int  // 顶点A
	nodeIndexB int  // 顶点B
	weight     int  // 权重
	isSelected bool // 是否选择
}

//采用邻接矩阵存储map
type Gmap struct {
	nodes     []Node // 数组存储结点列表
	matrix    []int  // 数组存储边数据矩阵
	capacity  int    // 图的容量
	nodeCount int    // 已添加顶点的个数
	edges     []Edge // 保存最小生成数的边
}

// 添加顶点
func (g *Gmap) AddNode(node Node) bool {
	if g.nodeCount >= g.capacity {
		return false
	}

	//g.nodes = append(g.nodes, node)	//此处有bug
	g.nodes[g.nodeCount] = node
	g.nodeCount++
	return true
}

// 重置所有的顶点到初始状态
func (g *Gmap) ResetNode() {
	for i := 0; i < g.nodeCount; i++ {
		g.nodes[i].isAccess = false
	}
}

//设置无向图的边
func (g *Gmap) SetUnGraphValue(row int, col int, weight int) bool {
	if row < 0 || row >= g.capacity || col < 0 || col >= g.capacity {
		return false
	}

	g.matrix[row*g.capacity+col] = weight
	g.matrix[col*g.capacity+row] = weight
	return true
}

// 输出矩阵
func (g *Gmap) Dump() {
	for i := 0; i < g.capacity; i++ {
		for j := 0; j < g.capacity; j++ {
			fmt.Print(g.matrix[i*g.capacity+j], "  ")
		}
		fmt.Println()
	}
}

// 深度优先遍历
func (g *Gmap) DepthFirstTraverse(startIndex int) {
	if startIndex < 0 || startIndex > g.capacity {
		return
	}

	if g.nodeCount != g.capacity { // 图并不完整
		return
	}

	fmt.Print(g.nodes[startIndex].value, "(", startIndex, ")")
	//设置该顶点已经访问过了
	g.nodes[startIndex].isAccess = true
	for i := 0; i < g.capacity; i++ {
		if g.matrix[i*g.capacity+startIndex] != 0 && g.nodes[i].isAccess != true { // 找到一个满足条件第一个边，然后递归搜索下去
			g.DepthFirstTraverse(i)
		}
	}

	return
}

// 广度优先遍历(核心是要实现分层)
func (g *Gmap) BreadthFirstTraverse(startIndex int) {
	if startIndex < 0 || startIndex > g.capacity {
		return
	}

	if g.nodeCount != g.capacity { // 图并不完整
		return
	}

	fmt.Print(g.nodes[startIndex].value, "(", startIndex, ")")
	//设置该顶点已经访问过了
	g.nodes[startIndex].isAccess = true
	// 保存符合条件的同一层的顶点索引
	var nodeIndexs []int
	nodeIndexs = make([]int, 0, g.capacity)

	for i := 0; i < g.capacity; i++ {
		if g.matrix[i*g.capacity+startIndex] != 0 && g.nodes[i].isAccess != true { // 找到一个满足条件第一个边，放入集合，先后继续找
			nodeIndexs = append(nodeIndexs, i)
			g.nodes[i].isAccess = true
			fmt.Print(g.nodes[i].value, "(", i, ")")
		}
	}

	if len(nodeIndexs) > 0 {
		g.breadthFirstTraverseImpl(nodeIndexs)
	}
	return
}

func (g *Gmap) breadthFirstTraverseImpl(nodeIndexs []int) {
	// 保存符合条件的同一层的顶点索引
	var curIndexs []int
	curIndexs = make([]int, 0, g.capacity)
	for i := 0; i < len(nodeIndexs); i++ {
		for j := 0; j < g.capacity; j++ {
			if g.matrix[j*g.capacity+nodeIndexs[i]] != 0 && g.nodes[j].isAccess != true { // 找到一个满足条件第一个边，放入集合，先后继续找
				curIndexs = append(curIndexs, j)
				g.nodes[j].isAccess = true
				fmt.Print(g.nodes[j].value, "(", j, ")")
			}
		}
	}

	if len(curIndexs) == 0 {
		return
	}

	g.breadthFirstTraverseImpl(curIndexs)
}

// 普利姆最小生成数算法
func (g *Gmap) PrimTree(startIndex int) {
	// 获取结点所有的边，找到最小的边，然后以这条边的另一个顶点为基础找最小的边，依次下去，直到找到的边数=顶点数-1就完毕。
	var (
		nodeIndex    int
		edges        []Edge // 待选边集合
		edge         Edge
		minEdgeIndex int
	)
	nodeIndex = startIndex
	edges = make([]Edge, 0, 10)
	g.nodes[nodeIndex].isAccess = true

	for {
		if len(g.edges) == g.capacity-1 {
			break
		}
		for i := 0; i < g.capacity; i++ {
			if g.matrix[i*g.capacity+nodeIndex] != 0 && g.nodes[i].isAccess != true { // 找到一个满足条件第一个边，放入集合，先后继续找
				edges = append(edges, Edge{nodeIndex, i, g.matrix[i*g.capacity+nodeIndex], false})

			}
		}

		// 返回一条最小的边索引
		minEdgeIndex = g.getMinEdge(edges)
		if minEdgeIndex == -1 {
			break
		}

		edge = edges[minEdgeIndex]
		fmt.Println(edge.nodeIndexA, "====", edge.nodeIndexB, "(", edge.weight, ")")
		edges[minEdgeIndex].isSelected = true
		g.edges = append(g.edges, edges[minEdgeIndex])

		//fmt.Println(g.edges)
		//fmt.Println(edges)
		// 下一次遍历这个顶点所有的边，因为是无向图，不能重复遍历，比如遍历A时AB边已经放入待选边，
		// 遍历B时，BA边就不能再放入待选边集合了
		nodeIndex = edge.nodeIndexB
		g.nodes[nodeIndex].isAccess = true

	}
}

// 获取最小的边
func (g *Gmap) getMinEdge(edges []Edge) (index int) {
	// TODO(tfbrother)如果没有找到最小边，一定要返回-1。因为默认值0是有意义的索引。花了半天时间才找到这个bug，哭
	index = -1
	var edge, minEdge Edge

	for i := 0; i < len(edges); i++ {
		edge = edges[i]
		if edge.isSelected == true {
			continue
		}

		if minEdge.weight == 0 {
			minEdge = edge
			index = i
		} else if minEdge.weight > edge.weight {
			minEdge = edge
			index = i
		}
	}

	return
}

type nodeSet []int

// 克鲁斯卡尔最小生成树算法
func (g *Gmap) KruskalTree() {
	var (
		edges           []Edge // 保存所有的边集合
		edgeSelectCount int
		minEdgeIndex    int
		sets            []nodeSet
		edgesSelected   []Edge
	)

	minEdgeIndex = -1
	edges = make([]Edge, 0, 2*g.capacity)
	sets = make([]nodeSet, 0, g.capacity)
	edgesSelected = make([]Edge, 0, g.capacity)

	// 1.取出所有的边
	for i := 0; i < g.capacity; i++ {
		for j := i + 1; j < g.capacity; j++ {
			if g.matrix[i*g.capacity+j] != 0 {
				edges = append(edges, Edge{i, j, g.matrix[i*g.capacity+j], false})
			}
		}
	}
	//fmt.Println(edges)

	for {
		if edgeSelectCount == g.capacity-1 {
			break
		}

		// 2.获取最小边
		minEdgeIndex = g.getMinEdge(edges)
		if minEdgeIndex == -1 {
			break
		}

		edges[minEdgeIndex].isSelected = true
		// 3.找出连接最小边的点，以及点所在的集合
		nodeA := edges[minEdgeIndex].nodeIndexA
		nodeB := edges[minEdgeIndex].nodeIndexB

		// 4.找出最小所在点的集合
		labelA := g.isInSet(sets, nodeA)
		labelB := g.isInSet(sets, nodeB)

		// 5.根据集合的不同，进行不同处理
		if labelA == -1 && labelB == -1 {
			sets = append(sets, []int{nodeA, nodeB})
		} else if labelA != -1 && labelB == -1 { // 把B点放入A点所在集合
			sets[labelA] = append(sets[labelA], nodeB)
		} else if labelA == -1 && labelB != -1 { // 把A点放入B点所在集合
			sets[labelB] = append(sets[labelB], nodeA)
		} else if labelA != 1 && labelB != 1 && labelA != labelB { // 合并两个集合
			sets[labelA] = append(sets[labelA], sets[labelB]...)
			// 切片中删除一个元素sets[labelB]
			sets = append(sets[:labelB], sets[labelB+1:]...)
		} else if labelA != 1 && labelB != 1 && labelA == labelB {
			continue
		}

		edgesSelected = append(edgesSelected, edges[minEdgeIndex])
		edgeSelectCount++
	}
	fmt.Println(edgesSelected)
	return
}

func (g *Gmap) isInSet(sets []nodeSet, nodeIndex int) (index int) {
	index = -1
	for i := 0; i < len(sets)-1; i++ {
		for j := 0; j < len(sets[i]); j++ {
			if sets[i][j] == nodeIndex {
				index = i
				return
			}
		}
	}

	return
}

// 初始化一个gmap
func NewGmap(capacity int) *Gmap {
	return &Gmap{
		nodes:     make([]Node, capacity),
		matrix:    make([]int, capacity*capacity),
		capacity:  capacity,
		nodeCount: 0,
		edges:     make([]Edge, 0, capacity-1),
	}
}