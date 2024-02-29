/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var res = make([][]int, 0)
	sort.Ints(nums)
	var i, start, end int
	for i = 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		start = i + 1
		end = len(nums) - 1
		for start < end {
			if nums[i]+nums[start]+nums[end] == 0 {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				start++
				for start < end && nums[start] == nums[start-1] {
					start++
				}
			} else if nums[i]+nums[start]+nums[end] < 0 {
				start++
			} else if nums[i]+nums[start]+nums[end] > 0 {
				end--
			}
		}
	}
	return res
}

// @lc code=end
