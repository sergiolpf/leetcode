package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var dummy *ListNode

	nodeVal := make(map[int]*ListNode)

	dummy = head

	for dummy.Next != nil {
		_, ok := nodeVal[dummy.Val]
		if ok {
			return dummy
		}

		nodeVal[dummy.Val] = dummy
		dummy = dummy.Next
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func main() {
	var node1, node2, node3, node4, node5, node6, node7 *ListNode
	node1 = &ListNode{3, &ListNode{}}
	node2 = &ListNode{2, &ListNode{}}
	node3 = &ListNode{0, &ListNode{}}
	node4 = &ListNode{-4, &ListNode{}}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	node5 = &ListNode{1, &ListNode{}}
	node6 = &ListNode{2, &ListNode{}}

	node5.Next = node6
	node6.Next = node5

	node7 = &ListNode{1, nil}

	//fmt.Printf("%v", detectCycle(node1))
	fmt.Printf("%v", detectCycle(node7))
}
