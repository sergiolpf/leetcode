package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList3(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	reversedHead := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedHead
}

func reverseList2(head *ListNode) *ListNode {
	var previous *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

func reverseList(head *ListNode) *ListNode {
	reversedList := ListNode{}
	dummy := &reversedList

	localNode := head
	nodes := []ListNode{}

	first := true

	for {
		if localNode == nil {
			if first {
				return nil
			}
			break
		}
		first = false
		nodes = append(nodes, *localNode)

		localNode = localNode.Next

	}

	if len(nodes) == 1 {
		return head
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		dummy.Val = nodes[i].Val
		if i == 0 {
			dummy.Next = nil
		} else {
			dummy.Next = &ListNode{nodes[i-1].Val, &ListNode{}}
			dummy = dummy.Next
		}
	}

	return &reversedList
}

func main() {
	fmt.Printf("%v", reverseList(&ListNode{}))
	fmt.Printf("%v", reverseList2(&ListNode{}))
	fmt.Printf("%v", reverseList3(&ListNode{}))
}
