/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
package leetcode

import "math"

func reverse(x int) int {
	var remainder, res int
	for x != 0 {
		//-231 <= x <= 231 - 1
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		remainder = x % 10
		x = x / 10
		res = res*10 + remainder
	}
	return res
}

// @lc code=end
