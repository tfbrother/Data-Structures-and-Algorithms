package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/algorithms/sorting"
	"github.com/tfbrother/Data-Structures-and-Algorithms/util"
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

	arr1 := util.GenrateRandomArray(10, 20, 50)
	fmt.Println(arr1)
	sorting.InsertionSort(arr1)
	fmt.Println(arr1)

	arr2 := util.GenrateRandomArray(10, 20, 50)
	fmt.Println(arr2)
	sorting.InsertionSortNew(arr2)
	fmt.Println(arr2)

	arr3 := util.GenrateRandomArray(10, 20, 50)
	fmt.Println(arr3)
	arr4 := sorting.MergeSort(arr3)
	fmt.Println(arr4)

	arr5 := util.GenrateRandomArray(10, 20, 50)
	fmt.Println(arr5)
	sorting.QuickSort1(arr5)
	fmt.Println(arr5)

	arr6 := util.GenrateRandomArray(30, 1, 2)
	fmt.Println(arr6)
	sorting.QuickSort3(arr6)
	fmt.Println(arr6)
}
