package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(n int) *ListNode {
	head := &ListNode{}
	node := head
	for i := 1; i <= n; i++ {
		node.Val = i
		node.Next = nil

		if i < n {
			node.Next = &ListNode{
				Val:  i + 1,
				Next: nil,
			}
		}
		node = node.Next
	}

	return head
}

func printNode(node *ListNode) {
	for {
		fmt.Printf("%d ", node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}
	fmt.Println()
}

// 位图
type bitMap struct {
	bits []byte
	cap  int
}

func newBitMap(cap int) *bitMap {
	bits := make([]byte, (cap>>3)+1)
	return &bitMap{bits: bits, cap: cap}
}

func (b *bitMap) add(num uint) {
	b.bits[num >> 3] |= 1 << (num & 0x07)
}

func (b *bitMap) isExist(num uint) bool {
	return b.bits[num >> 3]&(1<<(num & 0x07)) != 0
}

func (b *bitMap) remove(num uint) {
	b.bits[num >> 3] &= ^(1 << (num & 0x07))
}

// func main()  {
// 	printNode(test1(createNode(8)))
// }

// 整个反转
func test1(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}