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
