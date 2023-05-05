package main

// 94 binary-tree-inorder-traversal 2022-11-21 15:05:28

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
//
// 	var res []int
//
// 	inorder(root, &res)
//
// 	return res
// }
//
// func inorder(root *TreeNode, res *[]int) {
// 	if root == nil {
// 		return
// 	}
//
// 	inorder(root.Left, res)
// 	*res = append(*res, root.Val)
// 	inorder(root.Right, res)
// }
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var node = root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		node = node.Right
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
