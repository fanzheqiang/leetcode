/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package leetcode

import "sort"

// @lc code=start
func twoSum(nums []int, target int) []int {
	copyNum := make([]int, len(nums))
	copy(copyNum, nums)
	a, b := getTwoNum(nums, target)
	var indexA, indexB int
	var selectA bool
	for index, value := range copyNum {
		if value == a && !selectA {
			indexA = index
			selectA = true
			continue
		}
		if value == b {
			indexB = index
		}
	}
	return []int{indexA, indexB}
}

func getTwoNum(nums []int, target int) (int, int) {
	sort.Ints(nums)
	var i, j = 0, len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == target {
			return nums[i], nums[j]
		}
		if nums[i]+nums[j] > target {
			j--
		}
		if nums[i]+nums[j] < target {
			i++
		}
	}
	return 0, 0
}

// @lc code=end
