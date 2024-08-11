package listnode

import "testing"

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/

func ExchangeTwoNode(head *ListNode) *ListNode {
	// 首先要知道，如果想操作“当前节点”，那么我们最好是从当前节点的前一个节点来进行操作。交换也是如此。
	// 所以虚拟头节点需要全程在被交换的两个节点前，来完成这两个节点的交换操作。
	// 同时还要注意保存交换后可能无法访问到的节点
	// 使用虚拟头节点。
	dummy := &ListNode{Next: head}
	pre := dummy // pre 指向这一轮交换的两个节点中,第一个节点的前一个节点
	for pre.Next != nil && pre.Next.Next != nil {
		// tmp是这轮交换的一个节点,保存一下，以便访问第二个节点
		tmp := pre.Next
		// 保存第三个节点，因为第二个节点和第一个节点交换后，就访问不到第三个了
		tmp1 := pre.Next.Next.Next
		// 交换  要注意，这个过程中不能使用tmp,因为可能会出现环。 所有的操作都在pre上进行，否则太绕了
		pre.Next = pre.Next.Next
		pre.Next.Next = tmp
		pre.Next.Next.Next = tmp1
		pre = pre.Next.Next
	}
	return dummy.Next
}

func TestExchange(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	result := ExchangeTwoNode(head)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
