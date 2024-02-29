/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] N 字形变换
 */

// @lc code=start
package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res []byte
	for i := 0; i < numRows; i++ {
		var k = i
		for k < len(s) {
			res = append(res, s[k])
			if i != 0 && i != numRows-1 && k+(2*(numRows-1-i)) < len(s) {
				res = append(res, s[k+(2*(numRows-1-i))])
			}
			k += (2 * (numRows - 1))
		}
	}
	return string(res)
}

// @lc code=end
