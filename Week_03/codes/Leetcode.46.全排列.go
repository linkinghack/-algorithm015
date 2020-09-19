/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	visited := map[int]bool{}
	result := [][]int{}
	comb := []int{}

	var dfs func()

	dfs = func() {
		if len(comb) == len(nums) {
			result = append(result, append([]int{}, comb...))
			return
		}

		for i, _ := range nums {
			if !visited[i] {
				comb = append(comb, nums[i])
				visited[i] = true

				dfs()

				comb = comb[:len(comb)-1]
				visited[i] = false
			}
		}
	}

	dfs()
	return result
}

// @lc code=end

