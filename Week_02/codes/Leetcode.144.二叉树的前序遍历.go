package codes

import "container/list"

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	traversalWithStack(root, &result)
	return result
}

// 方法一 递归
func traversalRecursively(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		traversalRecursively(root.Left, result)
		traversalRecursively(root.Right, result)
	}
}

// 方法二 利用栈
func traversalWithStack(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	stack := list.New()
	var curr *TreeNode = root

	stack.PushBack(root)
	for stack.Len() > 0 {
		// 先根遍历
		// 每个节点先把右子树压栈, 最后访问
		top := stack.Back()
		stack.Remove(top)
		curr = top.Value.(*TreeNode)
		*result = append(*result, curr.Val) // 访问节点

		if curr.Right != nil {
			stack.PushBack(curr.Right)
		}
		if curr.Left != nil {
			stack.PushBack(curr.Left)
		}
	}
}

// @lc code=end
