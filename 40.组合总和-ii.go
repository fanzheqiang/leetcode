/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
package leetcode

import "sort"

var res = make([][]int, 0)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res = make([][]int, 0)
	getTargetNum(candidates, target, 0, []int{}, false)
	return res
}

func getTargetNum(candidates []int, target int, idx int, selectNum []int, selectCurrent bool) {
	if target == 0 {
		res = append(res, append([]int{}, selectNum...))
		return
	}
	if idx == len(candidates) || target < 0 {
		return
	}
	getTargetNum(candidates, target, idx+1, selectNum, false)

	if idx > 0 && candidates[idx-1] == candidates[idx] && !selectCurrent {
		return
	}
	selectNum = append(selectNum, candidates[idx])
	getTargetNum(candidates, target-candidates[idx], idx+1, selectNum, true)
	selectNum = selectNum[:len(selectNum)-1]
}

// @lc code=end
