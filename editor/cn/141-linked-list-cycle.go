package main

import (
	"fmt"
)

// 141 linked-list-cycle 2022-11-28 15:01:34

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	h := createNode(2)
	printNode(h)
	fmt.Println(hasCycle(h))
}
