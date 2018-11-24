package sorting

// 归并排序

func MergeSort(arr []int) (ret []int) {
	if len(arr) == 1 {
		return arr
	}
	middle := len(arr) / 2
	arr1 := MergeSort(arr[0:middle])
	arr2 := MergeSort(arr[middle:])
	ret = merge(arr1, arr2)
	return
}

func merge(arr1 []int, arr2 []int) (ret []int) {
	ret = make([]int, len(arr1)+len(arr2))
	var j, k int

	for i := 0; i < len(ret); i++ {
		if j > len(arr1)-1 { // arr1已经遍历完了或者
			ret[i] = arr2[k]
			k++
		} else if k > len(arr2)-1 { //arr2已经遍历完了或者
			ret[i] = arr1[j]
			j++
		} else if arr1[j] >= arr2[k] {
			ret[i] = arr2[k]
			k++
		} else if arr1[j] < arr2[k] {
			ret[i] = arr1[j]
			j++
		}
	}

	return
}
