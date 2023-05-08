package main

// 226 invert-binary-tree 2023-05-07 22:52:46

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	invertTree226(root)

	return root
}

func invertTree226(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
