package listnode

import "testing"

/*
	环形链表
给你一个链表的头节点 head ，判断链表中是否有环。
*/

// 时间复杂度：O(n) 空间复杂度：O(1)
func hasCycle(head *ListNode) bool {
	// 使用快慢指针，快指针每次走两步，慢指针每次走一步
	// 如果链表中有环，那么快指针最终会追上慢指针
	// 如果链表中没有环，那么快指针会先到达链表的尾部
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func TestHasCycle(t *testing.T) {
	// 测试用例
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	//head.Next.Next.Next.Next = head.Next.Next
	result := hasCycle(head)
	t.Log(result)
}
