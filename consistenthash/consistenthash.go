// 一致性hash算法
package consistenthash

import (
	"hash/crc32"
	"log"
	"sort"
	"strconv"
)

type Hash func([]byte) uint32

type Map struct {
	hash     Hash        //hash计算函数
	replicas int         // replicas是指的是每个节点的虚拟节点的个数。
	nodes    []int       // 保存所有的节点（实际节点和虚拟节点，排序了的，排序形成一个hash圆环）
	hashMap  map[int]int //保存虚拟节点对应的实际节点
}

// 添加节点
func (m *Map) Add(nodes []int) (err error) {
	for _, key := range nodes {
		// 将key转换成string
		keyStr := strconv.Itoa(int(key))
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(keyStr + string(i))))
			m.hashMap[hash] = int(key)
			m.nodes = append(m.nodes, hash)
		}
	}

	sort.Ints(m.nodes)
	return
}

// 根据key返回所在的缓存节点
func (m *Map) Get(key string) (node int) {
	hash := int(m.hash([]byte(key)))

	for _, v := range m.nodes {
		if v >= hash {
			return m.hashMap[v]
		}
	}

	return m.hashMap[m.nodes[0]]
}

func (m *Map) Dump() {
	log.Println(m.nodes, m.hashMap)
}

func New(replicas int) *Map {
	return &Map{
		replicas: replicas,
		hash:     crc32.ChecksumIEEE,
		hashMap:  make(map[int]int),
	}
}
