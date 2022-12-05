package main

// 86 partition-list 2022-10-25 17:07:50

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	smallHead, bigHead := new(ListNode), new(ListNode)
	small, big := smallHead, bigHead

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		}else {
			big.Next = head
			big = big.Next
		}

		head = head.Next
	}

	big.Next = nil

	small.Next = bigHead.Next

	return smallHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
