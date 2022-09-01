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
