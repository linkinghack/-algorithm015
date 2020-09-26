package codes

func numIslands(grid [][]byte) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	height := len(grid)
	width := len(grid[0])
	var dfsFlat func(i, j int)
	// 深度优先的将i,j 对应点连接的所有1标记为0
	dfsFlat = func(i, j int) {
		// 递归终止：触碰到边界或已不是陆地(!=1)
		if i < 0 || j < 0 || i >= height || j >= width || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		// 深度优先搜索领结的点
		dfsFlat(i+1, j)
		dfsFlat(i, j+1)
		dfsFlat(i-1, j)
		dfsFlat(i, j-1)
	}

	count := 0
	// 遍历grid, 标记岛屿数量
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				count++
				dfsFlat(i, j)
			}
		}
	}
	return count
}
