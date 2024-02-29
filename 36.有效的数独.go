/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
package leetcode

func isValidSudoku(board [][]byte) bool {
	var numberMap = [10][][2]int{}
	for line, linebytes := range board {
		for row, item := range linebytes {
			if item == '.' {
				continue
			}
			if !checkRepeat(numberMap[item-'0'], [2]int{line, row}) {
				return false
			}
			numberMap[item-'0'] = append(numberMap[item-'0'], [2]int{line, row})
		}
	}
	return true
}
func checkRepeat(existMap [][2]int, position [2]int) bool {
	for _, item := range existMap {
		if item[0] == position[0] || item[1] == position[1] {
			return false
		}
		if item[0]/3 == position[0]/3 && item[1]/3 == position[1]/3 {
			return false
		}
	}
	return true
}

// @lc code=end
