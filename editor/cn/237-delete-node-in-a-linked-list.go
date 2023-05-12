package main

// 237 delete-node-in-a-linked-list 2023-05-08 17:26:15

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	// if node.Next != nil { // 题目意思是该节点不会是tail，所以不需要判断
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	// }
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
