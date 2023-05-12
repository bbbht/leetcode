package main

// 543 diameter-of-binary-tree 2023-05-12 14:01:32

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	diameterOfBinaryTree543(root, &res)

	return res
}

func diameterOfBinaryTree543(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := diameterOfBinaryTree543(root.Left, res)
	right := diameterOfBinaryTree543(root.Right, res)

	if left+right > *res {
		*res = left + right
	}

	return max543(left, right) + 1
}

func max543(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
