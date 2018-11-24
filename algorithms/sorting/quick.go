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
	for k := 1; k < len(arr); k++ {
		if arr[k] < arr[0] {
			arr[i+1], arr[k] = arr[k], arr[i+1]
			i++
		}
	}
	arr[0], arr[i] = arr[i], arr[0]

	return i
}

// 针对近似有序数组排序的优化版版本1
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

// 针对有很多重复元素情况的优化版本
func QuickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}

	p := partition1(arr)
	QuickSort1(arr[0:p])
	QuickSort1(arr[p+1:])

	return
}

// 采用双路快速排序法，将相等的元素近似平均分部在左右两个部分
func partition2(arr []int) (i int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int() % len(arr)

	arr[0], arr[r] = arr[r], arr[0]
	i, j := 1, len(arr)-1

	for {
		for i < len(arr) && arr[i] < arr[0] {
			i++
		}

		for j > i && arr[j] > arr[0] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[0], arr[j] = arr[j], arr[0]
	return j
}
