/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
package leetcode

func longestCommonPrefix(strs []string) string {
	var start, end = 0, len(strs)
	if start == end-1 {
		return strs[start]
	}
	if start+1 == end-1 {
		return getTwoCommonPrefix(strs[start], strs[end-1])
	}
	var mid = (start + end) / 2
	s1 := longestCommonPrefix(strs[start:mid])
	s2 := longestCommonPrefix(strs[mid:end])

	return getTwoCommonPrefix(s1, s2)
}

func getTwoCommonPrefix(s1, s2 string) string {
	var l, s = s1, s2
	if len(l) < len(s) {
		l, s = s, l
	}
	for index, value := range []byte(s) {
		if l[index] != value {
			return s[:index]
		}
	}
	return s
}

// @lc code=end
