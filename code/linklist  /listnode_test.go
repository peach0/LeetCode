package linklist

import (
	"testing"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{7,nil}
	l2 := ListNode{8,nil}
	l3 := addTwoNumbers(&l1,&l2)
	t.Error(l3.Val)
}
//2. 两数相加
//本来就是逆序，从头个十百千万，过十进一，添加一个变量记录那个一，
//有点麻烦的是尾节点总是要创建一个，会存在空的现象，所以每次先创建一个节点存数据。
//最后返回节点的next节点。执行时间4ms,内存消耗5MB
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node1 := l1
	node2 := l2
	l3 := new(ListNode)
	node3 := l3
	a := 0
	for (node1 != nil || node2 != nil || a > 0) {
		node3.Next = new(ListNode)
		node3 = node3.Next
		b := 0
		c := 0
		if node1 != nil {
			b = node1.Val
		}
		if node2 != nil {
			c = node2.Val
		}
		node3.Val = (a + b + c) % 10
		a = (a + b + c) / 10
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}
	return l3.Next
}

func TestMergeList(t *testing.T) {
	var lists []*ListNode
	lists = append(lists,&ListNode{1, nil})
	lists[0].Next = &ListNode{4, nil}
	lists[0].Next.Next = &ListNode{5, nil}
	lists = append(lists,&ListNode{1, nil})
	lists[1].Next = &ListNode{3, nil}
	lists[1].Next.Next = &ListNode{4, nil}
	lists = append(lists,&ListNode{2, nil})
	lists[2].Next = &ListNode{6, nil}
	t.Log(mergeKLists(lists))
}

//23. 合并K个排序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	list :=  getMergeList(lists, 0 , len(lists) -1)
	return list
}
func getMergeList(lists[] *ListNode, start int, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if end - start == 1 {
		return compare(lists[start], lists[end])
	}
	l1 :=getMergeList(lists, start, (start + end) / 2)
	l2 :=getMergeList(lists,(start + end) / 2 + 1 , end)
	//return compare(getMergeList(lists, start, (start + end) / 2), getMergeList(lists,(start + end) / 2 + 1 , end))
	return compare(l1,l2)
}

func compare(node1 *ListNode, node2 *ListNode) *ListNode {
	switch {
	case node1 == nil && node2 == nil:
		return node2
	case node1 == nil:
		return node2
	case node2 == nil:
		return node1
	}
	var p ,l2 *ListNode
	if node1.Val <= node2.Val {
		p = node1
		l2 = node2
	}else {
		p = node2
		l2 = node1
	}

	pre := p
	l1 := p.Next
	for l1 != nil || l2 != nil  {
		if l1== nil {
			pre.Next = l2
			break
		}
		if l2 == nil {
			break
		}
		if l1.Val < l2.Val{
			pre = pre.Next
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
			pre = pre.Next
			pre.Next = l1
		}
	}
	return p
}
