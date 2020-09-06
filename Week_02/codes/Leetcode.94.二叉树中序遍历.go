package codes

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	traversalRecursively(root, &result)
	return result
}

// 方法一 递归
func traversalRecursively(root *TreeNode, result *[]int) {
	if root != nil {
		traversalRecursively(root.Left, result)
		*result = append(*result, root.Val)
		traversalRecursively(root.Right, result)
	}
}

// 方法二 利用栈
func traversalWithStack(root *TreeNode, result *[]int) {
	stack := list.New()
	var curr *TreeNode = root

	for curr != nil || stack.Len() > 0 {
		// 中序遍历，每个节点都先遍历左子树
		// 当前节点入栈，先访问左子树
		// 入栈顺序保证了弹出访问是中序遍历
		for curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		}
		top := stack.Back()
		stack.Remove(top)
		curr = top.Value.(*TreeNode)
		*result = append(*result, curr.Val) // 访问节点
		curr = curr.Right
	}
}
