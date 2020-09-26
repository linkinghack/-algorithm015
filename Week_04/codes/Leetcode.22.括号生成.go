/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package codes

// @lc code=start
func generateParenthesis(n int) []string {
	result := []string{}
	var dfs func(left, right int, current string)
	dfs = func(left, right int, current string) {
		if left == n && right == n {
			result = append(result, current)
			return
		}

		// 左括号不超总量总可以添加
		if left < n {
			dfs(left+1, right, current+"(")
		}

		// 右括号添加要考虑合法性，不超过左括号已添加数量就可以
		if right < left {
			dfs(left, right+1, current+")")
		}
	}

	dfs(0, 0, "")
	return result
}

// @lc code=end
