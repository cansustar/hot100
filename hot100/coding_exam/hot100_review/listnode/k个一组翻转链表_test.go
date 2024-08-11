package listnode

import (
	"fmt"
	"testing"
)

/*
 k个一组翻转链表
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
*/

// 时间复杂度：O(n) 空间复杂度：O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 可以使用递归法。分别找到每一组的头和尾，然后翻转这一组
	// 然后将翻转后的头和尾连接到上一组的尾和下一组的头
	// 递归的终止条件是，剩余节点不足k个，所以需要一个count来表示剩余节点的个数

	count := 0
	cur := head
	// 移动到该组的末尾的下一个节点. 因为要翻转的是head到cur之间的节点，所以cur是该组的末尾的下一个节点
	for count < k && cur != nil {
		cur = cur.Next
		count++
	}
	// 如果该组的数量等于k,翻转该组
	if count == k {
		// 翻转该组， reverse_head就是翻转后的头节点
		reverseHead := Reverse(head, cur)
		// 递归翻转下一组，然后将翻转后的头节点连接到该组的尾节点的下一个节点
		// 传入的cur就是下一组的第一个节点
		// head经过上面的Reverse,已经变成了该组的尾节点.由他来连接下一组
		fmt.Println("head:", head.Val)
		head.Next = reverseKGroup(cur, k)
		return reverseHead
	} else {
		// 如果剩余节点不足k个，那么直接返回head
		return head
	}
}

// 翻转链表
func Reverse(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != end {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	fmt.Println(start.Val)
	// 而返回的pre就是翻转后的这一组的头节点
	return pre
}

func TestReverseK(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	k := 3
	result := reverseKGroup(head, k)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
