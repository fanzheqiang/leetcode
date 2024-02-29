/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
package leetcode

func searchInsert(nums []int, target int) int {
	return getIndex(nums, 0, len(nums)-1, target)
}
func getIndex(nums []int, start, end, target int) int {
	if start > end {
		return start
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if target < nums[mid] {
		return getIndex(nums, start, mid-1, target)
	} else {
		return getIndex(nums, mid+1, end, target)
	}
}

// @lc code=end
