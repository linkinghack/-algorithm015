package codes

func updateBoard(board [][]byte, click []int) [][]byte {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		var mCount byte = 0
		for i := 0; i < 8; i++ {
			tx := x + dx[i]
			ty := y + dy[i]
			// 搜索领进8个点中的地雷数量
			if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) {
				continue
			}
			if board[tx][ty] == 'M' {
				mCount++
			}
		}

		if mCount > 0 {
			board[x][y] = '0' + mCount // 注意保存字符
		} else {
			board[x][y] = 'B'
			// dfs
			for i := 0; i < 8; i++ {
				tx := x + dx[i]
				ty := y + dy[i]
				if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) || board[tx][ty] == 'B' {
					// 递归终止
					continue
				}
				dfs(tx, ty)
			}
		}
	}

	cx, cy := click[0], click[1]
	if board[cx][cy] == 'M' {
		// 点到地雷
		board[cx][cy] = 'X'
	} else {
		dfs(cx, cy)
	}
	return board
}
