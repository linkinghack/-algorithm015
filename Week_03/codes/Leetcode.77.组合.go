/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := [][]int{}
	dfs(n, k, 1, []int{}, &result)
	return result
}

func dfs(n int, k int, start int, comb []int, result *[][]int) {
	if len(comb) == k {
		*result = append(*result, comb[0:])
		return
	}

	for i := start; i <= n-(k-len(comb))+1; i++ {
		tempcomb := make([]int, len(comb))
		copy(tempcomb, comb)
		dfs(n, k, i+1, append(tempcomb, i), result)
	}
}

// @lc code=end

