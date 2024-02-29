/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
package leetcode

func letterCombinations(digits string) []string {
	digitsNumMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	if len(digits) == 0 {
		return []string{}
	}

	res := []string{""}
	var tmp = make([]string, 0)
	for _, num := range []byte(digits) {
		digs := digitsNumMap[byte(num)]
		for _, dig := range []byte(digs) {
			for _, item := range res {
				tmp = append(tmp, item+string(dig))
			}
		}

		res = tmp
		tmp = make([]string, 0)
	}
	return res
}

// @lc code=end
