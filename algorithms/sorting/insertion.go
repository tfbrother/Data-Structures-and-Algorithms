package sorting

// 插入排序
// 思想：把数组分成两个部分，已经排序部分和未排序部分，每次遍历就是从未排序的里面依次拿一个元素出来放在已经排序部分的合适位置。
// 从头部开始的话，第一个元素肯定就是已经排序的

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

	return
}

// 优化版：思路就是减少内存交换的次数
// 1，4，5，2 比如1，4，5是有序的，找到2的位置时才进行插入，而不是让2和5交换位置，然后2再和4交位置。
func InsertionSortNew(arr []int) {
	for i := 1; i < len(arr); i++ {
		var j int
		a := arr[i]
		for j = i; j > 0; j-- {
			if a < arr[j-1] {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = a
	}

	return
}
