package main

// 142 linked-list-cycle-ii 2023-05-06 20:44:58

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	cycle := false
	fast, slow := head, head

	for slow != nil && fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			cycle = true
			break
		}
	}

	if !cycle {
		return nil
	}

	// fast从头开始，利用步数相同在环处相遇
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}

	return fast
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
