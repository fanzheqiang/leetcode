/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
package leetcode

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	var num, lastNum = 0, str[0]
	var res string
	for _, value := range []byte(str) {
		if value == lastNum {
			num++
			continue
		}
		res = res + strconv.Itoa(num) + string(lastNum)
		num = 1
		lastNum = value
	}
	res = res + strconv.Itoa(num) + string(lastNum)
	return res
}

// @lc code=end
