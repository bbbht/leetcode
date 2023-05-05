package main

// 101 symmetric-tree 2023-05-04 15:50:49

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs101(root.Left, root.Right)
}

func dfs101(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return dfs101(left.Left, right.Right) && dfs101(left.Right, right.Left)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
