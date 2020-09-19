/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
import "sort"

func permuteUnique(nums []int) [][]int {
	visited := map[int]bool{}
	result := [][]int{}
	comb := []int{}

	// 涉及到内容去重,依赖排列
	sort.Ints(nums)

	var dfs func()

	dfs = func() {
		if len(comb) == len(nums) {
			result = append(result, append([]int{}, comb...))
			return
		}

		for i, _ := range nums {
			if visited[i] {
				continue
			}

			// 消除重复组合
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			comb = append(comb, nums[i])
			visited[i] = true

			dfs()

			comb = comb[:len(comb)-1]
			visited[i] = false
		}
	}

	dfs()

	return result
}

// @lc code=end

