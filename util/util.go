package util

import (
	"math/rand"
	"time"
)

// 生成随机数数组
func GenrateRandomArray(num int, start int, end int) (ret []int) {
	ret = make([]int, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		value := rand.Int()%(end-start+1) + start
		ret[i] = value
	}

	return
}

// 生成三角形数组triangle array
func GenrateTriangleArray(num int, start int, end int) (ret [][]int) {
	ret = make([][]int, num)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		ret[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			value := rand.Int()%(end-start+1) + start
			ret[i][j] = value
		}
	}

	return
}
