package main

// 114 flatten-binary-tree-to-linked-list 2023-05-08 19:19:21

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var list []int
	preorder114(root, &list)

	cur := root
	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{Val: list[i], Left: nil, Right: nil}
		cur = cur.Right
	}
	return
}

func preorder114(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preorder114(root.Left, res)
		preorder114(root.Right, res)
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
