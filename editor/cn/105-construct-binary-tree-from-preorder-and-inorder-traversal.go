package main

// 105 construct-binary-tree-from-preorder-and-inorder-traversal 2023-05-06 16:36:57

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTree105(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildTree105(preorder []int, pLeft, pRight int, inorder []int, iLeft, iRight int) *TreeNode {
	if pLeft > pRight || iLeft > iRight {
		return nil
	}

	i := 0
	// 先序的首个节点为根节点
	for i = iLeft; i <= iRight; i++ {
		if preorder[pLeft] == inorder[i] {
			break
		}
	}
	cur := &TreeNode{Val: preorder[pLeft]}
	cur.Left = buildTree105(preorder, pLeft+1, pLeft+i-iLeft, inorder, iLeft, i-1)
	cur.Right = buildTree105(preorder, pLeft+i-iLeft+1, pRight, inorder, i+1, iRight)

	return cur
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
