/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
package leetcode

func nextPermutation(nums []int) {
	var start int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			start = i + 1
			break
		}
	}

	var p, q = start, len(nums) - 1
	for p < q {
		nums[p], nums[q] = nums[q], nums[p]
		p++
		q--
	}

	if start != 0 {
		for j := start; j < len(nums); j++ {
			if nums[j] > nums[start-1] {
				nums[start-1], nums[j] = nums[j], nums[start-1]
				break
			}
		}
	}
}

// @lc code=end
