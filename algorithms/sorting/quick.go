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
	QuickSort2(arr[0:p])
	QuickSort2(arr[p+1:])

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

// 针对有很多重复元素情况的优化版本
// 三路快速排序算法
func QuickSort3(arr []int) {
	if len(arr) <= 1 {
		return
	}

	i, j := partition3(arr)
	QuickSort3(arr[0:i])
	QuickSort3(arr[j:])

	return
}

// 采用三路快速排序法，将相等的元素近似平均分部在左右两个部分
func partition3(arr []int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int() % len(arr)

	arr[0], arr[r] = arr[r], arr[0]
	i, j, k := 0, 1, len(arr)

	for j < k {
		if arr[j] < arr[0] {
			arr[i+1], arr[j] = arr[j], arr[i+1]
			i++
			j++
		} else if arr[j] > arr[0] {
			arr[k-1], arr[j] = arr[j], arr[k-1]
			k--
		} else if arr[j] == arr[0] {
			j++
		}
	}

	arr[0], arr[i] = arr[i], arr[0]
	return i, j
}

// 三路快速排序算法改进版
func QuickSort4(arr []int) {
	if len(arr) > 12 {
		i, j := partition4(arr)
		QuickSort4(arr[0:i])
		QuickSort4(arr[j:])
	}

	if len(arr) > 1 { // 数据量较小时，采用插入排序更优
		InsertionSortNew(arr)
	}
	return
}

// Tukey’s ninthe
// Tukey’s ninther 方法选择切分元素：选择三组，每组三个元素，分别取三组元素的中位数，然后取三个中位数的中位数作为切分元素。
func partition4(arr []int) (int, int) {
	n := len(arr)
	m := n / 2
	//Tukey’s ninthe
	if n > 40 {
		s := n / 8
		medianOfThree(arr, 0, s, 2*s)
		medianOfThree(arr, m, m-s, m+s)
		medianOfThree(arr, n-1, n-1-s, n-1-2*s)
	}
	medianOfThree(arr, 0, m, n-1)

	i, j, k := 0, 1, n

	for j < k {
		if arr[j] < arr[0] {
			arr[i+1], arr[j] = arr[j], arr[i+1]
			i++
			j++
		} else if arr[j] > arr[0] {
			arr[k-1], arr[j] = arr[j], arr[k-1]
			k--
		} else if arr[j] == arr[0] {
			j++
		}
	}

	arr[0], arr[i] = arr[i], arr[0]
	return i, j
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
// @ 这个设计非常巧妙，就是参数m0在中间哦
func medianOfThree(arr []int, m1, m0, m2 int) {
	// sort 3 elements
	if arr[m1] < arr[m0] {
		arr[m0], arr[m1] = arr[m1], arr[m0]
	}

	if arr[m2] < arr[m1] {
		arr[m1], arr[m2] = arr[m2], arr[m1]
		if arr[m1] < arr[m0] {
			arr[m0], arr[m1] = arr[m1], arr[m0]
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

// IsSorted reports whether data is sorted.
func IsSorted(data []int) bool {
	n := len(data)
	for i := n - 1; i > 0; i-- {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
