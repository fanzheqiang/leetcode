/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var num []int
	for x != 0 {
		num = append(num, x%10)
		x = x / 10
	}

	var start, end = 0, len(num) - 1
	for start <= end {
		if num[start] != num[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// @lc code=end
