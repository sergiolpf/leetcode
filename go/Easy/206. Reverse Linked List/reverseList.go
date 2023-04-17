package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
}
