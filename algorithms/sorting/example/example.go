package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/algorithms/sorting"
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

	arr := []int{1, 4, 5, 7, 2, 3, 8, 10, 20, 8, 7}
	sorting.SelectSort(arr)

	fmt.Println(arr)

}
