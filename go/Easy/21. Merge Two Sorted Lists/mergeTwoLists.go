package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := ListNode{}
	myNode := &mergedList
	first := true

	for {
		if list1 == nil && list2 == nil {
			myNode = nil
			if first {
				return nil
			}
			break
		}
		first = false

		if list1 == nil && list2 != nil {
			myNode.Val = list2.Val
			myNode.Next = list2.Next
			break
		}
		if list2 == nil && list1 != nil {
			myNode.Val = list1.Val
			myNode.Next = list1.Next
			break
		}

		if list1 == nil && list2 == nil {
			break
		}

		if list1.Val == list2.Val {
			value := list1.Val
			localNode := &ListNode{}

			list1 = list1.Next
			list2 = list2.Next

			myNode.Val = value

			if list1 == nil && list2 == nil {
				myNode.Next = &ListNode{value, nil}
			} else {
				myNode.Next = &ListNode{value, localNode}
			}

			myNode = myNode.Next.Next

			continue
		}

		if list1.Val < list2.Val {
			myNode.Val = list1.Val
			myNode.Next = &ListNode{}
			myNode = myNode.Next

			list1 = list1.Next
			continue
		}

		if list1.Val > list2.Val {
			myNode.Val = list2.Val
			myNode.Next = &ListNode{}
			myNode = myNode.Next

			list2 = list2.Next
		}
	}
	return &mergedList
}
func main() {

	fmt.Println("*v", mergeTwoLists(&ListNode{1, nil}, &ListNode{2, nil}))
}
