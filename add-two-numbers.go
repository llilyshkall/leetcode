package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret1 := l1
	ret2 := l2
	overflow := 0
	tmp := 0
	for l1.Next != nil && l2.Next != nil {
		tmp = l1.Val + l2.Val + overflow
		l1.Val = tmp
		l2.Val = tmp
		overflow = tmp / 10
		l1.Val %= 10
		l2.Val %= 10
		l1 = l1.Next
		l2 = l2.Next
	}
	ret := ret1
	l := l1
	if l1.Next == nil {
		ret = ret2
		l = l2
	}
	tmp = l1.Val + l2.Val + overflow
	l.Val = tmp
	overflow = tmp / 10
	l.Val %= 10
	for overflow == 1 {
		if l.Next == nil {
			l.Next = &ListNode{Val: 1}
			overflow = 0
		} else {
			l = l.Next
			l.Val += overflow
			overflow = l.Val / 10
			l.Val %= 10
		}
	}
	return ret
}
