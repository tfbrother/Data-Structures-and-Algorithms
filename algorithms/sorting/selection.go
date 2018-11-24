package sorting

// 选择排序
// 第一次循环找到最小值的索引，然后把最小值放在第一个位置，第二次循环把最小值放在第二个位置
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}

	return
}
