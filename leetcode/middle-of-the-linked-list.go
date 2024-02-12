/*
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.
https://leetcode.com/problems/middle-of-the-linked-list/description/
INPUT:
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/
package leetcode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodeMap := make(map[int]*ListNode, 100)
	nodeSize := 0

	if head == nil {
		fmt.Printf("Empty list: returning.")
		return nil
	}
	if head.Next == nil {
		fmt.Printf("List of size 1, returning 1.")
		return head
	}

	for currentNode := head; currentNode != nil; currentNode = currentNode.Next {
		nodeMap[nodeSize] = currentNode
		fmt.Printf("nodeMap[%d] = %d at %p\n", nodeSize, currentNode.Val, currentNode)
		nodeSize++
	}
	fmt.Printf("List size= %d, half size + 1 = %d\n", nodeSize, nodeSize/2+1)
	return nodeMap[nodeSize/2]
}

func TestMiddleNode(t *testing.T) {
	var node6 ListNode = ListNode{Val: 6, Next: nil}
	var node5 ListNode = ListNode{Val: 5, Next: &node6}
	var node4 ListNode = ListNode{Val: 4, Next: &node5}
	var node3 ListNode = ListNode{Val: 3, Next: &node4}
	var node2 ListNode = ListNode{Val: 2, Next: &node3}
	var node1 ListNode = ListNode{Val: 1, Next: &node2}

	want := &node4
	got := middleNode(&node1)
	if got != want {
		t.Errorf("Got wrong node: val = %d instead of %d", got.Val, want.Val)
	}
}
