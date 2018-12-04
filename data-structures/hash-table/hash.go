package hash_table

import (
	"hash/crc32"
)

// 素数
const PRIME int = 98317

type HashMap map[string]string

type HashTable struct {
	count int       // 记录当前hash表中存放了多少个元素
	prime int       // hash值的素数
	data  []HashMap // 存放hash值的数据
}

func (h *HashTable) Add(key string, val string) {
	if !h.Contains(key) { //存在则更新，不存在则添加
		h.count++
	}

	h.put(key, val)
}

func (h *HashTable) Get(key string) string {
	if val, ok := h.data[h.hash(key)][key]; ok {
		return val
	}

	return ""
}

func (h *HashTable) hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key)) % uint32(h.prime)
}

// 判断是否存在key
func (h *HashTable) Contains(key string) bool {
	_, ok := h.data[h.hash(key)][key]
	return ok
}

func (h *HashTable) put(key string, val string) {
	h.data[h.hash(key)][key] = val
}

func NewHashTable() HashTable {
	hashTable := HashTable{
		count: 0,
		prime: PRIME,
		data:  make([]HashMap, PRIME),
	}

	for i := 0; i < PRIME; i++ {
		hashTable.data[i] = HashMap{}
	}

	return hashTable
}
