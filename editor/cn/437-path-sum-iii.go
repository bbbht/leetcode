package main

// 437 path-sum-iii 2023-05-08 17:04:05

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 计算以当前节点为根开始向下，符合要求的路径数
	res := pathSum437(root, targetSum)
	// 不包含当前节点，左右子树的符合要求数
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

func pathSum437(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if targetSum == root.Val {
		res++
	}

	res += pathSum437(root.Left, targetSum-root.Val)
	res += pathSum437(root.Right, targetSum-root.Val)

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
