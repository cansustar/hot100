package listnode

import (
	"fmt"
	"testing"
)

/*
合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/

func MergeTwoOrderList(list1 *ListNode, list2 *ListNode) *ListNode {
	// 使用一个哑节点作为新链表的头节点
	dummy := &ListNode{}
	cur := dummy
	// 遍历两个链表，比较两个链表的节点值，将较小的节点添加到新链表中
	for list1 != nil && list2 != nil {
		// 比较两个链表的节点值，将较小的那个连接到新链表中
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}
	// 如果其中一个链表遍历完了，另一个链表还有剩余节点，直接将剩余节点连接到新链表中
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

func TestMergeTwoOrderList(t *testing.T) {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 3}
	head1.Next.Next = &ListNode{Val: 7}
	head2 := &ListNode{Val: 0}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 4}
	result := MergeTwoOrderList(head1, head2)

	// 打印合并后的链表
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
