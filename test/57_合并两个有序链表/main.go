package main

import (
	"fmt"
)

func main() {

	//123
	//456
	//142536
	list := mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{4, &ListNode{5, &ListNode{6, nil}}})

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	tmpL1 := l1
	tmpL2 := l2
	newLHead := tmpL1
	if tmpL1.Val > tmpL2.Val {
		newLHead = tmpL2
		tmpL2 = tmpL2.Next
	} else {
		tmpL1 = tmpL1.Next
	}

	tmpHead := newLHead

	for {
		if tmpL1 == nil {
			tmpHead.Next = tmpL2
			return newLHead
		}

		if tmpL2 == nil {
			tmpHead.Next = tmpL1
			return newLHead
		}

		if tmpL1.Val <= tmpL2.Val {
			tmpHead.Next = tmpL1
			tmpL1 = tmpL1.Next
			tmpHead = tmpHead.Next
		} else {
			tmpHead.Next = tmpL2
			tmpL2 = tmpL2.Next
			tmpHead = tmpHead.Next
		}
	}
}
