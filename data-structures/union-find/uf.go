package union_find

/*
并查集
并查集是一种用来管理元素分组情况的数据结构。并查集可以高效地进行如下操作。不过需要注意并查集虽然可以进行合并操作，但是无法进行分割操作。
主要操作：
	1.查询元素a和元素b是否属于同一集合。
	2.合并元素a和元素b所在的集合。
	3.查询元素a在那一集合
	4.添加元素a
应用习题：https://www.cnblogs.com/douzujun/category/1003909.html
*/

type UnionFind struct {
	count  int
	parent []int
	rank   []int
	size   []int
}

// 查找过程, 查找元素a所对应的集合编号
// O(h)复杂度, h为树的高度
func (u *UnionFind) Find(a int) int {
	if a < 0 || a > u.count {
		return -1
	}

	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[a] == a
	for a != u.parent[a] {
		a = u.parent[a]
	}
	return a
}

// 判断元素a和元素b是否同属于一个集合
// O(h)复杂度, h为树的高度
func (u *UnionFind) IsConnected(a int, b int) bool {
	return u.Find(a) == u.Find(b)
}

// 合并元素a和元素b所在的集合。
func (u *UnionFind) Union(a int, b int) {
	aParentId, bParentId := u.Find(a), u.Find(b)
	if aParentId != bParentId {
		u.parent[aParentId] = bParentId //进行合并
	}
}

// TODO 基于size的优化=======================
// 合并元素a和元素b所在的集合。
func (u *UnionFind) UnionS(a int, b int) {
	aParentId, bParentId := u.Find(a), u.Find(b)
	if aParentId != bParentId {
		if u.size[aParentId] > u.size[bParentId] {
			u.parent[bParentId] = aParentId //进行合并
			u.size[aParentId] += 1
		} else {
			u.parent[aParentId] = bParentId //进行合并
			u.size[bParentId] += 1
		}

	}
}

// TODO 基于rank的优化=======================
func (u *UnionFind) FindR(a int) int {
	if a < 0 || a > u.count {
		return -1
	}

	// 不断去查询自己的父亲节点, 直到到达根节点
	for a != u.parent[a] {
		u.parent[a] = u.parent[u.parent[a]] // 路径压缩的代码，就这一行，每次压缩一个
		a = u.parent[a]
	}
	return a
}

// 合并元素a和元素b所在的集合。
func (u *UnionFind) UnionR(a int, b int) {
	aParentId, bParentId := u.Find(a), u.Find(b)
	if aParentId != bParentId {
		if u.rank[aParentId] > u.rank[bParentId] {
			u.parent[bParentId] = aParentId //进行合并
			u.rank[aParentId] += 1
		} else {
			u.parent[aParentId] = bParentId //进行合并
			u.rank[bParentId] += 1
		}

	}
}

func New(count int) *UnionFind {
	u := UnionFind{
		count:  count,
		parent: make([]int, count, count),
		rank:   make([]int, count, count),
		size:   make([]int, count, count),
	}

	for i := 0; i < count; i++ {
		u.parent[i] = i
		u.rank[i] = 1
		u.size[i] = 1
	}

	return &u
}
