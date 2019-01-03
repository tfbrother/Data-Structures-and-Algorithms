package main

// http://poj.org/problem?id=1163
// num triangle 数字三角形

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/util"
	"math"
)

const NUM = 100

var numTriangle [][]int
var maxNum [][]int = make([][]int, NUM)
var maxNum1 []int = make([]int, NUM)

func main() {
	numTriangle = util.GenrateTriangleArray(NUM, 10, 99)
	for i := 0; i < len(numTriangle); i++ {
		maxNum[i] = make([]int, NUM)
		maxNum1[i] = numTriangle[NUM-1][i]
	}
	//fmt.Println(maxNum1)
	fmt.Println(MaxSum(0, 0))
	fmt.Println(MaxSum1(0, 0))
}

// 自顶向下
func MaxSum(i, j int) int {
	// end case
	if maxNum[i][j] == 0 {
		if i == NUM-1 {
			maxNum[i][j] = numTriangle[i][j]
		} else {
			x := MaxSum(i+1, j)
			y := MaxSum(i+1, j+1)
			maxNum[i][j] = int(math.Max(float64(x), float64(y))) + numTriangle[i][j]
		}
	} // else {} has calculate 已经计算过了

	return maxNum[i][j]
}

// 自底向上，优化空间复杂度
func MaxSum1(i, j int) int {
	for a := NUM - 2; a >= i; a-- {
		for b := j; b <= a; b++ {
			//fmt.Println(b)
			maxNum1[b] = int(math.Max(float64(maxNum1[b]), float64(maxNum1[b+1]))) + numTriangle[a][b]
		}
	}

	return maxNum1[j]
}
