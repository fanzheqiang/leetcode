/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res = make([][]int, 0)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] || nums[i]+nums[length-3]+nums[length-2]+nums[length-1] < target {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[length-2]+nums[length-1] < target {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			var start, end = j + 1, length - 1
			for start < end {
				tmp := nums[i] + nums[j] + nums[start] + nums[end]
				if tmp == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					for start < end && nums[start] == nums[start-1] {
						start++
					}
				} else if tmp > target {
					end--
				} else {
					start++
				}
			}
		}
	}
	return res
}

// @lc code=end
