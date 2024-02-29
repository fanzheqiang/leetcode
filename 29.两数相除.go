/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
package leetcode

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	var smailThanZero bool
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		smailThanZero = true
	}
	var res int
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}

	var i, tmp int
	for dividend >= divisor {
		tmp = divisor
		i = 1
		for dividend >= tmp {
			if math.MaxInt32-i <= res {
				if !smailThanZero {
					return math.MaxInt32
				}
				if smailThanZero && math.MaxInt32-i+1 <= res {
					return math.MinInt32
				}
			}

			res += i
			dividend = dividend - tmp
			i = i << 1
			tmp = tmp << 1
		}

	}
	if smailThanZero {
		res = 0 - res
	}

	return res
}

// @lc code=end
