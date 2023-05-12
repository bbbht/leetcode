package main

// 337 house-robber-iii 2023-05-08 17:51:12

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	a, b := rob337(root)

	return max337(a,b)
}

func rob337(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}

	la, lb := rob337(root.Left)
	ra, rb := rob337(root.Right)

	// 分为两种
	// 不打当前节点，则左右都可打
	a = max337(la, lb) + max337(ra, rb)
	// 打当前，则左右不可打,而是加上其不打本节点的返回值
	b = root.Val + la + ra

	return
}

func max337(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
