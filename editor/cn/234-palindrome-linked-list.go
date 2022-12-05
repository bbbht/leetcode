package main

// 234 palindrome-linked-list 2022-12-01 15:01:00

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var pre *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		// 反转前半部分
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}

	for pre != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
