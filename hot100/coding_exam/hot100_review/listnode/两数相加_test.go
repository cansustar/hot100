package listnode

import (
	"fmt"
	"testing"
)

/*
两数相加
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
eg:
	输入：l1 = [2,4,3], l2 = [5,6,4]
	输出：[7,0,8]
*/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		关键在于，进位的处理.
		每次都要计算进位carry=sum / 10
		新链表的值为sum % 10
	*/
	// 初始化进位为0
	carry := 0
	// 定义一个哑节点作为新链表的头节点
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		// 初始化两个链表中当前值
		x, y := 0, 0
		// 因为两个链表的长度可能不同，所以需要判断链表是否遍历完
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		// 计算当前节点的值
		sum := x + y + carry
		// 计算进位
		carry = sum / 10
		// 创建新节点
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	// 如果最后一位相加后有进位，需要在新链表中添加一个节点
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

func TestTwoSum(t *testing.T) {
	list1 := &ListNode{Val: 2}
	list1.Next = &ListNode{Val: 7}
	list1.Next.Next = &ListNode{Val: 8}

	list2 := &ListNode{Val: 4}
	list2.Next = &ListNode{Val: 9}
	list2.Next.Next = &ListNode{Val: 7}
	list2.Next.Next.Next = &ListNode{Val: 4}
	result := AddTwoNumbers(list1, list2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
