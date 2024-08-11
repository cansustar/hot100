package listnode

/*
环形链表2, 需要返回链表开始入环的第一个节点
*/

func hasCycleAndFindNode(head *ListNode) *ListNode {
	// 同样是快慢指针，
	//当相遇时，将慢指针（或者快指针也行）指向head,然后快慢指针同时走，
	//再次相遇的地方就是入环的第一个节点
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
