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
// 缺点：对于序列中重复值不多的情况下，它比传统的2分区快速排序需要更多的交换次数。
func QuickSort3(arr []int) {
	if len(arr) <= 1 {
		return
	}

	i, j := partition3(arr)
	QuickSort3(arr[0:i])
	QuickSort3(arr[j:])

	return
}

// 采用三路快速排序法
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
// 三向切分法选择中轴点
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

// Quicksort5, following Bentley and McIlroy,
// ``Engineering a Sort Function,'' SP&E November 1993.
// 思路：在普通的三分区快速排序的基础上，对一般的快速排序进行了改进。在划分过程中，i遇到的与v相等的元素交换到最左边，
// j遇到的与v相等的元素交换到最右边，i与j相遇后再把数组两端与v相等的元素交换到中间
func QuickSort5(arr []int) {
	if len(arr) > 12 {
		i, j := partition5(arr)
		QuickSort5(arr[0:i])
		QuickSort5(arr[j:])
	}

	if len(arr) > 1 { // 数据量较小时，采用插入排序更优
		InsertionSortNew(arr)
	}
	return
}

func partition5(arr []int) (int, int) {
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

	return BentleyMcIlroyPartition(arr)
}

func ErrPartition5(arr []int) (int, int) {
	n := len(arr)

	// TODO 边界条件中存在bug
	a, b, c, d := 1, 1, n-1, n-1
	for {
		for b < c {
			if arr[b] < arr[0] {
				b++
			} else if arr[b] == arr[0] {
				arr[b], arr[a] = arr[a], arr[b]
				a++
				b++
			} else { // arr[b] >arr[0]
				break
			}
		}

		for b < c {
			if arr[c] > arr[0] {
				c--
			} else if arr[c-1] == arr[0] {
				arr[c], arr[d-1] = arr[d-1], arr[c]
				c--
				d--
			} else { // arr[b] >arr[1]
				break
			}
		}

		if b >= c {
			break
		}
		arr[b], arr[c] = arr[c], arr[b]
		b++
		c--
	}

	e := b
	for i := 0; i <= e; i++ {
		arr[e], arr[i] = arr[i], arr[e]
		e--
	}

	f := n - 1
	for i := c; i <= f; i++ {
		arr[f], arr[i] = arr[i], arr[f]
		f--
	}

	return b - a, n - d + c
}

// 主要用于测试QuickSort5的边界条件设置的是否合理
// 输入 [20 9 8 7 1]
// 输出 [1 7 8 9 20]
func ErrQuickSort5(arr []int) {
	if len(arr) > 1 {
		i, j := ErrPartition5(arr)
		ErrQuickSort5(arr[0:i])
		ErrQuickSort5(arr[j:])
	}

	return
}

//BentleyMcIlroyPartition
func BentleyMcIlroyPartition(arr []int) (int, int) {
	n := len(arr)
	// TODO 这种边界的设置，边界条件设置非常有技巧，下面的约束条件一定要理解清楚。
	// Invariants are:
	//	data[0] = pivot (set up by ChoosePivot)
	//	data[0 <= i < a] = pivot
	//	data[a <= i < b] < pivot
	//	data[b <= i < c] is unexamined
	//	data[c <= i < d] > pivot
	//	data[d <= i < n] = pivot
	// 1.a指向的前一个元素就是左边最后一个等于pivot的元素，所以初始时a=1，
	// 2.b指向的元素表示第一个待比较的元素。所以初始时a=1，结束时b可能指向的是一个不存在的元素，初始值设置为b=1
	// (此时一直从左往右比较交换，b最后一次指向的就是数组最后一个索引位置的后一个索引，越界)。
	// 此时data[a <= i < b]这个区间为空，data[a <= i < b]也为空
	// 3.c指向的是左边第一个大于pivot的元素，初始情况下，这个元素应该是不存在的。所以初始值应该设置为：c=n，此时data[c <= i < d]这个区间也为空。
	// 4.d指向的是右边第一个等于pivot的元素，初始情况下，这个元素应该也是不存在的。所以初始值应该设置为：d=n，此时data[d <= i < n]这个区间也为空。
	// 这样的初始值设置，都满足约束条件
	a, b, c, d := 1, 1, n, n
	// TODO 思考，这个循环跳出的时候，b和c之间的值是怎么样的？
	// b>=c 都有可能。
	for {
		// 假设我们整个数组里面arr[b] < arr[0] || arr[b] == arr[0]都成立，则只需要执行这一个for循环即完成了这组划分。
		// 如果c的初始值设置成c=n-1，那么arr[n-1] 将无法和 arr[0]比较。
		// 如果条件设置为b<n或者b<c+1，都会存在数组访问越界。
		for b < c {
			if arr[b] < arr[0] {
				b++
			} else if arr[b] == arr[0] {
				arr[b], arr[a] = arr[a], arr[b]
				a++
				b++
			} else { // arr[b] >arr[0]
				break
			}
		}

		for b < c {
			if arr[c-1] > arr[0] {
				c--
			} else if arr[c-1] == arr[0] {
				arr[c-1], arr[d-1] = arr[d-1], arr[c-1]
				c--
				d--
			} else { // arr[b] >arr[1]
				break
			}
		}

		if b >= c {
			break
		}
		arr[b], arr[c-1] = arr[c-1], arr[b]
		b++
		c--
	}

	e := b - 1
	for i := 0; i < a; i++ {
		arr[e], arr[i] = arr[i], arr[e]
		e--
	}

	f := n - 1
	for i := 0; i < n-d; i++ {
		arr[f], arr[i+c] = arr[i+c], arr[f]
		f--
	}

	return b - a, n - d + c
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
