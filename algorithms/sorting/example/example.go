package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/algorithms/sorting"
)

func main() {
	// 测试冒泡排序
	bubbleArr := []int{1, 8, 7, 5, 10, 32, 48}

	// 升序
	sorting.Bubble(bubbleArr, 0)
	fmt.Println(bubbleArr)

	// 降序
	sorting.Bubble(bubbleArr, 1)
	fmt.Println(bubbleArr)
}
