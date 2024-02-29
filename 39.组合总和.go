/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package leetcode

var res = make([][]int, 0)

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	getTargetNum(candidates, target, 0, []int{})
	return res
}

func getTargetNum(candidates []int, target int, idx int, selectNum []int) {
	if target == 0 {
		res = append(res, append([]int{}, selectNum...))
		return
	}
	if idx == len(candidates) || target < 0 {
		return
	}
	getTargetNum(candidates, target, idx+1, selectNum)

	selectNum = append(selectNum, candidates[idx])
	getTargetNum(candidates, target-candidates[idx], idx, selectNum)
	selectNum = selectNum[:len(selectNum)-1]
}

// @lc code=end
