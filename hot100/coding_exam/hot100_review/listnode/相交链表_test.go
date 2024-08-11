package listnode

import "testing"

/*
相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
eg:
	输入：listA = [4,1,8,4,5], listB = [5,0,1,8,4,5]
	输出：Intersected at '8'
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 两个指针分别从两个链表的头节点开始遍历，当一个指针到达链表的尾部时，将它重定向到另一个链表的头部
	// 这样两个指针就会同时遍历两个链表，且步长一致，当两个指针相等时，就是两个链表相交的节点

	// 解释： 想象两个指针在同一个链表的不同起点，想要在同一个节点相遇，那他们之间就需要消除相对距离。
	// 如何消除相对距离？ 就是让他们走的距离是相同的。
	//A指针走完链表A再走链表B，B指针走完链表B再走链表A，这样两个指针走的距离就是相等的。
	// 如果两个链表相交，那么两个指针就会在相交的节点相遇，如果两个链表不相交，那么两个指针就会在链表的尾部的空节点相遇
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}

func TestGetIntersectionNode(t *testing.T) {
	listA := &ListNode{Val: 4}
	listA.Next = &ListNode{Val: 1}
	listA.Next.Next = &ListNode{Val: 8}
	listA.Next.Next.Next = &ListNode{Val: 4}
	listA.Next.Next.Next.Next = &ListNode{Val: 5}

	listB := &ListNode{Val: 5}
	listB.Next = &ListNode{Val: 0}
	listB.Next.Next = &ListNode{Val: 1}
	listB.Next.Next.Next = listA.Next.Next

	result := getIntersectionNode(listA, listB)
	t.Log(result.Val)
}
