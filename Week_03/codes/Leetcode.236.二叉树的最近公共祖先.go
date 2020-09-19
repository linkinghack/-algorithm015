/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
package codes

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 左子树包含p或q
	right := lowestCommonAncestor(root.Right, p, q) // 右子树包含p或q

	if right != nil && left != nil {
		return root
	} else if right == nil {
		return left
	} else {
		return right
	}
}

// @lc code=end
