/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
package leetcode

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := i; j < len(haystack); j++ {
			if haystack[j] != needle[j-i] {
				break
			}

			if j-i == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

// @lc code=end
