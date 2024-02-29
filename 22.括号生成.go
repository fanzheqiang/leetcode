/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
package leetcode

var num int

func generateParenthesis(n int) []string {
	num = n
	return generate([]string{""}, n)
}

func generate(tmp []string, n int) []string {
	if len(tmp[0]) == 2*num {
		return tmp
	}
	var tmp1 []string
	if n > 0 {
		var left []string
		for _, item := range tmp {
			left = append(left, item+"(")
		}
		tmp1 = generate(left, n-1)
	}

	var tmp2 []string
	if 2*(num-n) > len(tmp[0]) {
		var right []string
		for _, item := range tmp {
			right = append(right, item+")")
		}
		tmp2 = generate(right, n)
	}

	return append(tmp1, tmp2...)
}

// @lc code=end
