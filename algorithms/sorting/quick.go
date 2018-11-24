package sorting

import (
	"math/rand"
	"time"
)

// 快速排序

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	p := partition(arr)
	QuickSort(arr[0:p])
	QuickSort(arr[p+1:])

	return
}

// 对arr进行partition操作
// 返回i，是的arr[0:i]<arr[i];arr[i+1]>arr[i]
func partition(arr []int) (i int) {
	val := arr[0]
	for k := 1; k < len(arr); k++ {
		if arr[k] < val {
			arr[i+1], arr[k] = arr[k], arr[i+1]
			i++
		}
	}
	arr[0], arr[i] = arr[i], arr[0]

	return i
}

// 优化版吧
func QuickSort1(arr []int) {
	if len(arr) <= 1 {
		return
	}

	p := partition1(arr)
	QuickSort1(arr[0:p])
	QuickSort1(arr[p+1:])

	return
}

// 主要针对的优化场景是，arr基本是已经排序的情况下，退化成链表排序了。
// 比如arr[0]已经是数组中最小的元素了，则分区每次就只能减少一个元素，近似退化成链表了，所以改成随机取一个索引来比较。
func partition1(arr []int) (i int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int() % len(arr)

	arr[0], arr[r] = arr[r], arr[0]
	for k := 1; k < len(arr); k++ {
		if arr[k] < arr[0] {
			arr[i+1], arr[k] = arr[k], arr[i+1]
			i++
		}
	}
	arr[0], arr[i] = arr[i], arr[0]

	return i
}
