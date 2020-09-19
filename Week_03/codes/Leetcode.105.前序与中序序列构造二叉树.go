/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIdx := map[int]int{}

	// 建立中序节点索引，遍历前序过程在中序序列中定位根节点位置
	for i, v := range inorder {
		inorderIdx[v] = i
	}

	var root *TreeNode
	var dfsBuildTree func(preLeft, preRight, inLeft, inRight int) *TreeNode
	dfsBuildTree = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		// terminator
		if preLeft > preRight {
			return nil
		}

		// process
		tempRoot := &TreeNode{Val: preorder[preLeft]}
		rootAtInorder := inorderIdx[preorder[preLeft]]
		leftSubTreeLength := rootAtInorder - inLeft
		rightSubTreeLength := inRight - rootAtInorder

		// left subtree
		tempRoot.Left = dfsBuildTree(preLeft+1, preLeft+leftSubTreeLength, inLeft, rootAtInorder-1)

		// right subtree
		tempRoot.Right = dfsBuildTree(preRight-rightSubTreeLength+1, preRight, rootAtInorder+1, inRight)

		return tempRoot
	}

	root = dfsBuildTree(0, len(preorder)-1, 0, len(inorder)-1)
	return root
}
