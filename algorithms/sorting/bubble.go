package sorting

// 冒泡排序 direction=0,表示升序，1表示降序
func Bubble(arr []int, direction int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if (direction == 0 && arr[i] < arr[j]) || (direction == 1 && arr[i] > arr[j]) { // 升序/降序
				arr[i], arr[j] = arr[j], arr[i] // go语言特有的一行代码交换值
			}
		}
	}
}

// 冒泡排序优化版 direction=0,表示升序，1表示降序
func BubbleUpgrade(arr []int, direction int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := false
		for j := i + 1; j < len(arr); j++ {
			if (direction == 0 && arr[i] < arr[j]) || (direction == 1 && arr[i] > arr[j]) { // 升序/降序
				arr[i], arr[j] = arr[j], arr[i] // go语言特有的一行代码交换值
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}
