package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/algorithms/sorting"
	"github.com/tfbrother/Data-Structures-and-Algorithms/algorithms/sorting/util"
)

func main() {
	// 升序
	bubbleArr := []int{1, 8, 7, 5, 10, 32, 48}
	sorting.Bubble(bubbleArr, 0)
	fmt.Println(bubbleArr)

	// 降序
	bubbleArr1 := []int{1, 8, 7, 5, 10, 32, 48}
	sorting.Bubble(bubbleArr1, 1)
	fmt.Println(bubbleArr1)

	// 优化版升序
	bubbleArr2 := []int{1, 8, 7, 5, 10, 32, 48}
	sorting.BubbleUpgrade(bubbleArr2, 0)
	fmt.Println(bubbleArr2)

	arr := util.GenrateRandomArray(10, 100, 200)
	fmt.Println(arr)
	sorting.SelectSort(arr)

	fmt.Println(arr)

}
