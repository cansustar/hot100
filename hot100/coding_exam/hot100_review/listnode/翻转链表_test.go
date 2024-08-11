package listnode

import "testing"

/*
翻转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
eg:
	输入：head = [1,2,3,4,5]
	输出：[5,4,3,2,1]
*/
// 时间复杂度：O(n) 空间复杂度：O(1)
func ReverseList(head *ListNode) *ListNode {
	// 改变指针方向即可，不要把这个和交换两个节点搞混了
	/* 使用迭代的方式。定义两个指针，一个指向当前节点，一个指向前一个节点
	在迭代的过程中，将当前节点的next指针指向前一个节点，然后两个指针都向后移动。
	最终当当前节点为nil时，前一个节点就是新的头节点
	在这个过程中需要注意的是，需要提前保存当前节点的下一个节点，
	因为在修改当前节点的next指针之后，当前节点的next指针原先指向的元素会丢失
	*/
	var pre *ListNode // 前一个节点
	cur := head       // 当前节点
	// 注意，一次循环里，保存当前节点的下一个节点，然后将当前节点的next指针指向前一个节点，最后两个指针向后移动
	for cur != nil {
		// 暂存当前节点的下一个节点
		tmp := cur.Next
		// 当前节点的next指针指向前一个节点
		cur.Next = pre
		// 两个指针向后移动:将pre移动到当前cur的位置，将cur移动到暂存的tmp的位置
		pre = cur
		cur = tmp
	}
	return pre
}

func TestReverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	result := ReverseList(head)
	t.Log(result.Val)
}
