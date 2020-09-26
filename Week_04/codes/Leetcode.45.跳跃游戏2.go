package codes

func jump(nums []int) int {
	furthest := 0 // 下一步最优的跳跃点
	steps := 0    // 统计跳跃步数
	standing := 0 // 目前所处位置

	// 贪心法解决，每次都选择跳跃到下一步最长的一点
	for i, v := range nums {
		// 如果刚好最后一步跳跃到目标位置（最大值没有超出去）
		//  会导致记录步数的if再进入一次，记录了多余的跳跃
		if i == len(nums)-1 {
			break
		}

		// 一趟循环，边解决跳跃，边解决下一步位置探测
		// 主要借助来standing变量来判断遍历过程中是否到达了之前已经决定跳跃的地点
		//  到达后将步数加上即可
		if v+i >= furthest {
			// 查找下一步跳跃点
			furthest = v + i
		}

		if i == standing {
			// 当前遍历idx：i 触碰到standing时表示在当前standing位置上能跳跃的点已全部探测到
			//   此时furthest记录的已是最优下一步
			standing = furthest // 下一步最优已确定
			steps++
		}
	}

	return steps
}
