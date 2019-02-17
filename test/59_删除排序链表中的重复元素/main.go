package main

import "fmt"

func main() {
	p := deleteDuplicates(&ListNode{1,&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{3,nil}}}}})

	for p != nil {
		fmt.Println(p.Val)
		p= p.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//1->1->2->3->3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p := head

	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next

		}else {
			p = p.Next
		}
	}
	return head
}
