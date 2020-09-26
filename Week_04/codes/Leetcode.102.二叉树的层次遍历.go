package codes

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		// 每层准备新的队列
		// 每次完成一层的遍历
		queueElems := []*TreeNode{}

		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				queueElems = append(queueElems, node.Left)
			}
			if node.Right != nil {
				queueElems = append(queueElems, node.Right)
			}
		}

		result = append(result, append([]int{}, level...))
		queue = queueElems
	}
	return result
}
