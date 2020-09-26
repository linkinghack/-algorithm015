/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package codes

// @lc code=start

func canJump(nums []int) bool {
	// 从末尾位置(目标),倒着推到回去
	standing := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// 每次进入判断是否可以有i位置跳跃到standing位置
		if nums[i]+i >= standing {
			// 发现i位置可跳跃到上一个检查点,
			// 后续只需要判断是否可以跳跃到此检查点即可
			standing = i
		}
	}

	if standing == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end
