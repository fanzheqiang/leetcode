/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package leetcode

func isValid(s string) bool {
	var systemMap = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var tmp []byte
	for _, value := range []byte(s) {
		item, ok := systemMap[value]
		if ok {
			if len(tmp) < 1 || tmp[len(tmp)-1] != item {
				return false
			}
			tmp = tmp[0 : len(tmp)-1]
			continue
		} else {
			tmp = append(tmp, value)
		}
	}

	if len(tmp) != 0 {
		return false
	}
	return true
}

// @lc code=end
