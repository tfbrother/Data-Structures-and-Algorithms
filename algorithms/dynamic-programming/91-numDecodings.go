package main

import "fmt"

/*
https://leetcode.com/problems/decode-ways/
动态规划
	版本：
	V1:
		Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode Ways.
		思路：此处的转移方程非常重要，要回溯前两个步骤。时间复杂度O(n)，空间复杂度O(n);可以考虑优化空间复杂度。
		状态转移方程如下：
			T(n),n=1,2 直接结算。
			T(n) , n > 2
			if  第n个数字无效，return 0
			else if 第n个数字只能单独作为一个编码
				T(n) = T(n-1)
			else if 第n个数字只能和第n-1个数字结合成一个编码
				T(n) = T(n-2)
			else	// n既能单独作为一个编码，也能和n-1结合
				T(n) = T(n-1) - T(n-2)
*/

// todo 字符为0时要特殊处理。
// 存在无效的数字时返回0，比如s中有两个连续的00，或者以0开头，或者0前面的那个数字大于2了。
func numDecodings(s string) int {
	var nums []int
	n := len(s)
	nums = make([]int, n+1)
	if s[0] == 48 {
		return 0
	}

	nums[0] = 1
	if n >= 2 {
		if s[1] == 48 {
			if s[0] > 50 {
				return 0
			} else {
				nums[1] = 1
			}
		} else if s[0] == 49 || (s[0] == 50 && s[1] < 55) {
			nums[1] = 2
		} else {
			nums[1] = 1
		}
	}

	for i := 2; i < n; i++ {
		// 0 的ascii码值为48
		if s[i] == 48 && (s[i-1] == 48 || s[i-1] > 50) { // 数字无效，连续出现两个0或者0前面的数字大于2
			return 0
		} else if s[i] == 48 { // i只能和i-1结合成一个编码
			nums[i] = nums[i-2]
		} else if s[i-1] > 50 || (s[i-1] == 50 && s[i] > 54) || s[i-1] == 48 { // i只能单独作为一个编码
			nums[i] = nums[i-1]
		} else { // i既能单独作为一个编码，也能和i-1结合
			nums[i] = nums[i-1] + nums[i-2]
		}
	}

	return nums[n-1]
}

func main() {
	s := "100"
	fmt.Println(numDecodings(s))
}
