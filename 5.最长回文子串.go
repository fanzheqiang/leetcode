/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package leetcode

// @lc code=start
func longestPalindrome(s string) string {
	var res, s1, s2, tmp string
	for i := 0; i < len(s); i++ {
		s1 = checkPalindrom(s, i, i)
		s2 = checkPalindrom(s, i-1, i)
		tmp = s2
		if len(s1) > len(s2) {
			tmp = s1
		}
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func checkPalindrom(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return s[left+1 : right]
}

// @lc code=end
