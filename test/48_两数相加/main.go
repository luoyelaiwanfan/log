package main

import "fmt"

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3,nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4,nil}}}

	newList := addTwoNumbers(l1, l2)
	for newList!= nil {
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}


//[2,4,3]
//[5,6,4]
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	newList := new(ListNode)
	p := newList
	p.Next = l1
	s1 := 0
	for {
		v1 :=0
		v2 := 0
		if l1 == nil && l2 == nil {
			if s1 != 0 {
				pre := new(ListNode)
				pre.Val = s1
				p.Next = pre
			}
			break
		}
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := s1 + v1 + v2


		if p.Next == nil {
			pre := new(ListNode)
			pre.Val = sum%10
			p.Next = pre

		}else {
			p.Next.Val = sum%10
		}
		p = p.Next
		s1 = sum/10

	}

	return newList.Next
}
