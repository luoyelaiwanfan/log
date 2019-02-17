package main

import "fmt"

func main() {
	//1->2->3->4->5
	head := new(ListNode)
	head.Val = 1

	p1 := new(ListNode)
	p1.Val = 2
	head.Next = p1

	p2 := new(ListNode)
	p2.Val = 3
	p1.Next = p2

	p3 := new(ListNode)
	p3.Val = 4
	p2.Next = p3

	//p4 := new(ListNode)
	//p4.Val = 5
	//p3.Next = p4

	//tmpHead := head
	//for tmpHead != nil {
	//	fmt.Println("v:", tmpHead.Val)
	//	tmpHead = tmpHead.Next
	//}

	reorderList(head)
	//test(head)
	for head != nil {
		fmt.Println("v:", head.Val)
		head = head.Next
	}

}


func test(head *ListNode) {
	head = nil
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Val  int
	Next *List
	Pre *List
}
//
//func reorderList(head *ListNode) {
//	if head == nil {
//		return
//	}
//	oldHead := head
//
//	h := new(List)
//	h.Val = head.Val
//	h.Pre = nil
//	h.Next = nil
//
//	t := h
//
//	// 编程双向链表
//	head = head.Next
//	for head != nil {
//		p := new(List)
//		p.Val = head.Val
//		p.Next=nil
//		p.Pre = t
//		t.Next = p
//		t = t.Next
//
//		head = head.Next
//	}
//
//	newHead := new(ListNode)
//	preHead := newHead
//	for h!=nil && t!=nil  {
//		if h == t {
//			 p := new(ListNode)
//			 p.Val = h.Val
//			 p.Next = nil
//			preHead.Next = p
//			preHead = preHead.Next
//			break
//		} else if h.Next == t{
//			p := new(ListNode)
//			p.Val = h.Val
//			p.Next = nil
//			preHead.Next = p
//			preHead = preHead.Next
//
//			p = new(ListNode)
//			p.Val = t.Val
//			p.Next = nil
//			preHead.Next = p
//			preHead = preHead.Next
//			break
//		}else {
//			p := new(ListNode)
//			p.Val = h.Val
//			p.Next = nil
//			preHead.Next = p
//			preHead = preHead.Next
//
//			p = new(ListNode)
//			p.Val = t.Val
//			p.Next = nil
//			preHead.Next = p
//			preHead = preHead.Next
//			h = h.Next
//			t = t.Pre
//		}
//	}
//	head = oldHead
//	if newHead.Next !=nil {
//		head.Val = newHead.Next.Val
//		head.Next = newHead.Next.Next
//	}
//
//}



func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//找到中间节点，快慢指针
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Println(slow)
	//翻转后半部分
	cur := slow.Next
	for cur.Next != nil {
		mov := cur.Next
		cur.Next = mov.Next
		mov.Next = slow.Next
		slow.Next = mov
	}
	//reorder
	p1, p2 := head, slow.Next
	for p1 != slow {
		slow.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = slow.Next
	}
}
