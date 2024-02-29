/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
package leetcode

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string
	var count int
	for index, i := range nums {
		count = num / i
		num = num % i
		for count > 0 {
			res += roman[index]
			count--
		}
	}
	return res
}

// @lc code=end
