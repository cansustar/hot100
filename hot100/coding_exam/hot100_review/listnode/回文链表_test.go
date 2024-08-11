package listnode

import "testing"

/*
	回文链表
	给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。
	如果是，返回 true ；否则，返回 false 。
*/
// 时间复杂度：O(n) 空间复杂度：O(1)
func isPalindList(head *ListNode) bool {
	// 如果一个链表是回文链表，那么后半部分的链表翻转后，应该和前半部分的链表一样
	// 要想翻转后半部分链表，那么首先需要找到链表的中点，然后将后半部分链表翻转
	// 链表中点可以使用快慢指针来找到
	middleNode := FindMiddleNode(head)
	// 翻转后半部分链表
	reverseNode := ReverseListNode(middleNode.Next)
	// 比较前半部分链表和翻转后的后半部分链表
	p1, p2 := head, reverseNode
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return true
}

// 找到链表中点
func FindMiddleNode(head *ListNode) *ListNode {
	// 使用快慢指针，快指针每次走两步，慢指针每次走一步。
	//当快指针走到链表尾部时，慢指针刚好走到链表的中点
	// 解释：为什么呢？
	// 想象链表长度变成两倍，那么走相同的次数，快指针到达末尾时，慢指针刚好到达中点
	// 当然要考虑可能链表长度是奇数，也可能是偶数。
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 翻转链表
func ReverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func TestIsPalindList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	result := isPalindList(head)
	t.Log(result)
}
