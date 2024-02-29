/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
package leetcode

func romanToInt(s string) int {
	var charMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res, num, lastNum int
	for _, value := range []byte(s) {
		num = charMap[value]
		if num > lastNum {
			res = res - 2*lastNum
		}
		res += num
		lastNum = num
	}
	return res
}

// @lc code=end
