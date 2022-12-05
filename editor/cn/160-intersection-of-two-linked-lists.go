package main

// 160 intersection-of-two-linked-lists 2022-11-28 16:34:50

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	sizeA, tailA, sizeB, tailB := 0, headA, 0, headB
	for tailA.Next != nil {
		sizeA++
		tailA = tailA.Next
	}
	for tailB.Next != nil {
		sizeB++
		tailB = tailB.Next
	}
	// 若相交，结尾定相同
	if tailB != tailA {
		return nil
	}

	for i := 0; i < sizeA-sizeB; i++ {
		headA = headA.Next
	}
	for i := 0; i < sizeB-sizeA; i++ {
		headB = headB.Next
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
