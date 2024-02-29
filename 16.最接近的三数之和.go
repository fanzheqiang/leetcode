/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res, i, start, end, gap int
	gap = math.MaxInt32
	for i = 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		start = i + 1
		end = len(nums) - 1
		for start < end {
			add := nums[i] + nums[start] + nums[end]
			if add-target == 0 {
				return target
			} else if add-target > 0 {
				end--
			} else {
				start++
			}

			tmp := add - target
			if tmp < 0 {
				tmp = -tmp
			}
			if tmp < gap {
				gap = tmp
				res = add
			}
		}
	}
	return res
}

// @lc code=end
