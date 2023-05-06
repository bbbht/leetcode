package main

// 236 lowest-common-ancestor-of-a-binary-tree 2023-05-06 09:31:36

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 递归遍历，直到遍历发现p、q，或者到了叶子节点，返回nil
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 分别在左右子树中，那么肯定就是当前节点是最近公共祖先
	if left != nil && right != nil {
		return root
	}
	// 否则只存在于左右一侧，那么返回的一侧节点就是结果
	if left != nil {
		return left
	}

	return right
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
