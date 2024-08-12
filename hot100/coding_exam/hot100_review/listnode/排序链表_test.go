package listnode

import "testing"

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

func OrderList(head *ListNode) *ListNode {
	/*
		之前做过，合并两个有序链表。
		所以可以使用归并的思想，将链表分成两部分，然后合并
	*/
	// 所以需要先找到链表的中点，使用快慢指针
	// 归并，实际上就是递归，所以需要先写递归的终止条件
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找到中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 分成两部分
	l1 := head
	l2 := slow.Next
	// 断开
	slow.Next = nil
	// 递归排序两个链表
	l1Sort := OrderList(l1)
	l2Sort := OrderList(l2)
	// 合并两个有序链表
	return MergeTwoList(l1Sort, l2Sort)

}

func MergeTwoList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func TestOrderList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 0}
	result := OrderList(head)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
