package sorting

import (
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/heap"
)

//堆排序
//思路就是用数组构造堆，然后依次弹出堆的根元素，就是排序好了的对象了。
func HeapSort(arr []int) {
	myHeap := heap.NewMaxHeap(len(arr))
	myHeap.Init(arr)

	for i := 0; i < len(arr); i++ {
		arr[i], _ = myHeap.Delete()
	}
	return
}
