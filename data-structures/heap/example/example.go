package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/heap"
	"math/rand"
	"strconv"
	"time"
)

type Element struct {
	key   int
	value string
}

func (e *Element) Less(item heap.Item) bool {
	return e.key < item.(*Element).key
}

func (e *Element) Swap(item heap.Item) {
	e.key, item.(*Element).key = item.(*Element).key, e.key
	e.value, item.(*Element).value = item.(*Element).value, e.value
}

func (e Element) ToString() string {
	return "key=" + strconv.Itoa(e.key) + ",value=" + e.value
}

func main() {
	// 堆
	myarr := GenrateRandomElements(30, 10, 100)
	//fmt.Println(myarr)
	myHeap := heap.NewMaxHeap(30)
	for _, v := range myarr {
		myHeap.Push(v)
	}
	//myHeap.Dump()
	for !myHeap.Empty() {
		item, err := myHeap.Pop()
		if err != nil {
			panic(err)
		}

		fmt.Print(item.(*Element).key, ":", item.(*Element).value, " ")
	}

	fmt.Println()

	myarr1 := GenrateRandomElements(30, 10, 100)
	myHeap.Init(myarr1)
	for !myHeap.Empty() {
		item, err := myHeap.Pop()
		if err != nil {
			panic(err)
		}

		fmt.Print(item.(*Element).key, ":", item.(*Element).value, " ")
	}

	fmt.Println()
}

// 生成随机数数组
// TODO 注意返回值必须是[]heap.Item类型，而不能是[]*Element类型。
func GenrateRandomElements(num int, start int, end int) (ret []heap.Item) {
	ret = make([]heap.Item, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		key := rand.Int()%(end-start+1) + start
		value := "val" + strconv.Itoa(key)
		ret[i] = &Element{key: key, value: value}
	}

	return
}
