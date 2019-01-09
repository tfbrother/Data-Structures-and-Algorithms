package main

/*
https://leetcode.com/problems/score-after-flipping-matrix/
贪心算法：翻转矩阵后的得分
	版本：
	V1:
		Runtime: 0 ms, faster than 100.00% of Go online submissions for Score After Flipping Matrix.
		思路：贪心策略。
			1.只有每行的第一个数字为0是才对该行进行变换
			2.每列的1的数量小于0的数量时进行"列变换"，此处在实现上并没有正在变换，只是计算时的一个技巧
		经过求幂的优化，执行时间由4ms减少到0ms。
*/

import (
	"fmt"
)

func matrixScore(A [][]int) int {
	row := len(A)
	column := len(A[0])
	var total int

	for i := 0; i < row; i++ {
		if A[i][0] == 0 { // 行变换，只要第一个元素是0就进行变换
			//inverse(A[i])
			// TODO 优化1
			for j := 0; j < column; j++ {
				A[i][j] = 1 - A[i][j]
			}
		}
		// 先把第一列的计算出来
		// 利用位运算计算来计算2的幂，不用math里面的pow
		// total += int(math.Pow(2, float64(column-1)))
		// TODO 优化2
		total += (1 << uint32((column - 1)))
	}

	for i := 1; i < column; i++ { // 统计每列中0的数量，只要大于1的数量，就进行翻转
		num1 := 0 // 记录每列中1/0的数量差
		num2 := 0 // 记录每列中1/0的数量差
		for j := 0; j < row; j++ {
			if A[j][i] == 1 {
				num1++
			} else {
				num2++
			}
		}

		if num1 > num2 { // 需要进行变换
			total += (1 << uint32((column - i - 1))) * num1
		} else {
			total += (1 << uint32((column - i - 1))) * num2
		}
	}

	return total
}

func inverse(a []int) {
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] == 1 {
			a[i] = 0
		} else {
			a[i] = 1
		}
	}
}

func main() {
	a := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}

	b := matrixScore(a)
	fmt.Println(b)
}
