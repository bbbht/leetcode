package main

import (
	"math"
)

// 124 binary-tree-maximum-path-sum 2023-05-06 15:47:22

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt

	maxPathSum124(root, &res)

	return res
}

func maxPathSum124(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	// 分别计算左右子树的最大值，小于0则忽略其路径
	left := max124(maxPathSum124(root.Left, res), 0)
	right := max124(maxPathSum124(root.Right, res), 0)

	*res = max124(*res, left+right+root.Val)

	// 只能选择一个路径
	return max124(left, right) + root.Val
}

func max124(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
