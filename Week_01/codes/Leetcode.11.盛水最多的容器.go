import "math"

// @lc code=start
func maxArea(height []int) int {
	var max int = 0
	i := 0
	j := len(height) - 1
	for i < j {
		area := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		if area > max {
			max = area
		}

		// if height[i] >= height[j] {
		// 	j--
		// } else {
		// 	i++
		// }

		// 多走几步的策略
		if height[i] >= height[j] {
			for i < j && height[j-1] <= height[j] {
				j--
			}
			j--
		} else {
			for i < j && height[i+1] <= height[i] {
				i++
			}
			i++
		}

	}

	return max
}
