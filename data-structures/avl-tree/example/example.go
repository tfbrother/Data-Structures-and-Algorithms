package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/avl-tree"
	"strconv"
)

type item struct {
	key   string
	value string
}

func (n *item) Less(b avl_tree.Item) bool {
	return n.key < b.(*item).key
}

func (n *item) ToString() string {
	return "key=" + n.key + ",value=" + n.value
}

func main() {
	avl := avl_tree.NewAVL()
	for i := 0; i < 100; i++ {
		n := &item{"tfbrother" + strconv.Itoa(i), strconv.Itoa(i)}
		avl.Add(n)
	}

	fmt.Println("树的大小", avl.Size())

	items := avl.InOrder()
	keys := make([]string, 0, len(items))

	for i := 0; i < len(items); i++ {
		keys = append(keys, items[i].(*item).key)
	}

	fmt.Println("所有的keys", keys)

	fmt.Println("是否是二叉搜索树", avl.IsBST())
	fmt.Println("是否是平衡二叉搜索树", avl.IsBalanced())
	fmt.Println("获取元素：tfbrother10，值为：", avl.Get(&item{"tfbrother10", "tfbrother"}))

	fmt.Println("删除元素tfbrother10测试=========")
	avl.Remove(&item{"tfbrother10", "tfbrother"})
	fmt.Println("删除后获取：tfbrother10，值为：", avl.Get(&item{"tfbrother10", "tfbrother"}))
	fmt.Println("删除元素后是否是平衡二叉搜索树", avl.IsBalanced())
}
