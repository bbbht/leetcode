package main
// 206 reverse-linked-list 2022-11-29 17:24:11

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}