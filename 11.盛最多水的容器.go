/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
package leetcode

func maxArea(height []int) int {
	var i, j = 0, len(height) - 1
	var res, tmp, minHeight int
	for i < j {
		minHeight = height[i]
		if height[j] < minHeight {
			minHeight = height[j]
		}
		tmp = minHeight * (j - i)
		if tmp > res {
			res = tmp
		}
		if height[j] < height[i] {
			j--
		} else {
			i++
		}
	}
	return res
}

// @lc code=end
