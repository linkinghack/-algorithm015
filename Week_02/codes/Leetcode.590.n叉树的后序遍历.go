package codes

import "container/list"

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}

	resultQueue := list.New()
	stack := list.New()
	var curr *Node
	// 前序遍历
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back()
		stack.Remove(top)
		curr = top.Value.(*Node)
		resultQueue.PushBack(curr) // 访问根
		// 子节点入栈
		for _, node := range curr.Children {
			stack.PushBack(node)
		}
	}

	// 反转顺序，变前序为后序
	for resultQueue.Len() > 0 {
		top := resultQueue.Back()
		resultQueue.Remove(top)
		result = append(result, top.Value.(*Node).Val)
	}

	return result
}

// @lc code=end
