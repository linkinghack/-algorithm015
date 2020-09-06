/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package codes

// @lc code=start
func twoSum(nums []int, target int) []int {
	numIdxMap := map[int]int{} // key -> num, value -> idx
	result := []int{0, 0}
	for i, v := range nums {
		wanted := target - v
		if wantedIdx, exists := numIdxMap[wanted]; exists {
			result[0] = wantedIdx
			result[1] = i
			break
		} else {
			numIdxMap[v] = i
		}
	}

	return result
}

// @lc code=end
