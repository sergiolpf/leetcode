package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	var node1, node2, node3, node4, node5, node6 *ListNode
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

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
