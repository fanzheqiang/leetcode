/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
package leetcode

func lengthOfLongestSubstring(s string) int {
	var start = 0
	var res int
	var strMap = make(map[byte]int)
	for index, value := range []byte(s) {
		if acc, ok := strMap[value]; ok && acc >= start {
			tmp := index - start
			if tmp > res {
				res = tmp
			}
			start = acc + 1
		}
		strMap[value] = index
	}
	tmp := len([]byte(s)) - start
	if tmp > res {
		res = tmp
	}
	return res
}

// @lc code=end
