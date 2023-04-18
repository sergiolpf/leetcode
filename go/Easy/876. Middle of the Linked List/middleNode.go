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

func main() {
	fmt.Printf("%v", middleNode(&ListNode{}))
}
