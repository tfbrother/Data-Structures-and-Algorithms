package sorting

// 冒泡排序 direction=0,表示升序，1表示降序
func Bubble(arr []int, direction int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if (direction == 0 && arr[i] < arr[j]) || (direction == 1 && arr[i] > arr[j]) { // 升序/降序
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}
