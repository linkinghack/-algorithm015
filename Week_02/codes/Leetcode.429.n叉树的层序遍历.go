package codes

import "container/list"

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)
	//result = append(result, []int{root.Val})
	for queue.Len() > 0 {
		layer := []int{}
		currentQueueLength := queue.Len() // 当前长度包括一整层的节点, 不做此处理在多叉树中无法得知本层节点
		// 不要求按层分开输出，仅要求输出层序遍历序列则不需要此length
		for i := 0; i < currentQueueLength; i++ {
			top := queue.Front()
			queue.Remove(top)
			node := top.Value.(*Node)
			layer = append(layer, node.Val)
			// 子树入队
			for _, subNode := range node.Children {
				queue.PushBack(subNode)
			}
		}

		if len(layer) > 0 {
			result = append(result, layer)
		}
	}
	return result
}

// @lc code=end
