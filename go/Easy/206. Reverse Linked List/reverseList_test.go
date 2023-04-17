package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test 1", args{&ListNode{1, nil}}, &ListNode{1, nil}},
		{"test 2", args{&ListNode{1, &ListNode{2, nil}}}, &ListNode{2, &ListNode{1, nil}}},
		{"test 3", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}},
		{"test 3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test 1", args{&ListNode{1, nil}}, &ListNode{1, nil}},
		{"test 2", args{&ListNode{1, &ListNode{2, nil}}}, &ListNode{2, &ListNode{1, nil}}},
		{"test 3", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}},
		{"test 3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList3(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{"test 1", args{&ListNode{1, nil}}, &ListNode{1, nil}},
		//{"test 2", args{&ListNode{1, &ListNode{2, nil}}}, &ListNode{2, &ListNode{1, nil}}},
		{"test 3", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}},
		//{"test 3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList3() = %v, want %v", got, tt.want)
			}
		})
	}
}
