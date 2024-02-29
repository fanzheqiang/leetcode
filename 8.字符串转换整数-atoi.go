/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	var res int
	var symbol int
	var start bool
	for _, value := range []byte(s) {
		if value == ' ' && !start {
			continue
		}
		if value == '-' && !start {
			symbol = 2
			start = true
			continue
		}
		if value == '+' && !start {
			symbol = 1
			start = true
			continue
		}
		if !(value >= '0' && value <= '9') {
			if symbol == 2 {
				return 0 - res
			}
			return res
		}
		if res > math.MaxInt32/10 {
			if symbol == 2 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		if res == math.MaxInt32/10 {
			if symbol != 2 && int(value-'0') > 7 {
				return math.MaxInt32
			}
			if symbol == 2 && int(value-'0') > 8 {
				return math.MinInt32
			}
		}
		res = res*10 + int(value-'0')
		start = true
	}
	if symbol == 2 {
		return 0 - res
	}
	return res
}

// @lc code=end
