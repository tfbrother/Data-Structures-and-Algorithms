package sorting

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
