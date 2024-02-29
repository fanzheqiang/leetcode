/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var index int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// @lc code=end
