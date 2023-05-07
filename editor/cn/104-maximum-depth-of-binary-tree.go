package main

// 104 maximum-depth-of-binary-tree 2023-05-06 16:30:26

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max104(left, right) + 1

}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
