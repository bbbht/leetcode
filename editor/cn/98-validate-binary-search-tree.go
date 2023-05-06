package main

import (
	"math"
)

// 98 validate-binary-search-tree 2023-05-06 15:37:51

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValid98(root, math.MaxInt, math.MinInt)
}

func isValid98(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	return root.Val < max && root.Val > min &&
		isValid98(root.Left, root.Val, min) && isValid98(root.Right, max, root.Val)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
