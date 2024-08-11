package listnode

import "testing"

/*
 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
eg:
	输入：head = [1,2,3,4,5], n = 2
	输出：[1,2,3,5]
*/

func DeleteDescNode(head *ListNode, n int) *ListNode {
	// 首先要确定，该如何找到倒数第n个节点。
	// 使用双指针，快指针先走n步，然后快慢指针一起走，当快指针到达链表尾部时，慢指针就是倒数第n个节点
	// 但需要注意， 如果要“删掉”倒数第n个，我们最后走到的是倒数第n+1个，所以我们需要找到倒数第n+1个节点
	// 删除，实际上就是把指向倒数第n个的next指针，指向倒数第n-1个节点
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	// 因为实际要让slow移动到倒数n+1个节点，所以fast要先移动n+1步（这样相比倒数n，slow就少移动一步）
	// 先让fast移动到n+1处
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	// 然后让fast和slow一起走，直到fast到达链表尾部，此时slow到达了倒数第n+1个节点（可以比手指数一数）
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func TestDelete(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	result := DeleteDescNode(head, 2)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
