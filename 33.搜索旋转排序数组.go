/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package leetcode

func search(nums []int, target int) int {
	return getIndex(nums, 0, len(nums)-1, target)

}

func getIndex(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[end] == target {
		return end
	}
	if nums[mid] > nums[start] {
		if target >= nums[start] && target < nums[mid] {
			return getIndex(nums, start, mid-1, target)
		} else {
			return getIndex(nums, mid+1, end, target)
		}
	} else {
		if target <= nums[end] && target > nums[mid] {
			return getIndex(nums, mid+1, end, target)
		} else {
			return getIndex(nums, start, mid-1, target)
		}
	}
}

// @lc code=end
