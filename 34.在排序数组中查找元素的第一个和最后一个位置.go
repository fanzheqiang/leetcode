/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package leetcode

func searchRange(nums []int, target int) []int {
	left := getIndex(nums, 0, len(nums)-1, target, true)
	right := getIndex(nums, 0, len(nums)-1, target, false)
	return []int{left, right}
}

func getIndex(nums []int, start, end, target int, left bool) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		if left && (mid == 0 || nums[mid-1] != target) {
			return mid
		}
		if !left && (mid == len(nums)-1 || nums[mid+1] != target) {
			return mid
		}
	}
	if target < nums[mid] || (left && target == nums[mid]) {
		return getIndex(nums, start, mid-1, target, left)
	} else {
		return getIndex(nums, mid+1, end, target, left)
	}
}

// @lc code=end
