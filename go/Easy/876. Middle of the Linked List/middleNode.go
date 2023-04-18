package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	var current *ListNode
	count := 0

	current = head

	for current != nil {
		count++
		current = current.Next
	}

	current = head
	for i := 1; i <= count/2; i++ {
		current = current.Next
	}

	return current

}

func middleNode2(head *ListNode) *ListNode {
	slow := head
	fast := head

	// move slow pointer one step at a time and fast pointer two steps at a time
	// when the fast pointer reaches the end of the list, the slow pointer will be at the middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	fmt.Printf("%v", middleNode(&ListNode{}))
	fmt.Printf("%v", middleNode2(&ListNode{}))
}
