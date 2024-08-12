package listnode

import "testing"

/*
	随机链表的复制
	给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
	要求返回这个链表的深拷贝。深拷贝应该正好由n个全新节点组成，其中每个新节点的值都设为其对应的原节点的值。
	新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，
	并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
	eg:
		输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
		输出：【7,null】【13,0】【11,4】【10,2】【1,0】
其实没啥区别
*/

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func RandomListCopy(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	// 难点其实在于，因为随机节点的存在，当我们复制一个节点时，对应的随机节点可能还没有创建，所以需要先创建随机节点，然后再连接
	// 但是随机节点，同样可能也指向其他创建或者未创建的节点
	// 迭代+节点拆分。
	// 1. 遍历原链表，复制每个节点，将复制的节点放在原节点的后面
	for node := head; node != nil; node = node.Next.Next {
		copyNode := &RandomListNode{Val: node.Val, Next: node.Next}
		node.Next = copyNode
	}
	// 2. 遍历原链表，复制随机节点
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			// 因为随机节点其实也是链表中的节点，所以随机节点也复制过了
			// 这里就是将原节点的随机节点的下一个节点赋值给复制节点的随机节点 Random.Next就是随机节点的复制节点
			node.Next.Random = node.Random.Next
		}
	}
	// 3. 拆分链表，即将复制节点从原链表中拆分出来，也就是让复制节点的next指向复制节点
	// 保存复制链表的头节点
	copyHead := head.Next
	for node := head; node != nil; node = node.Next {
		copyNode := node.Next
		// 让原节点的next指向原本的节点，也就是复制节点的next
		node.Next = copyNode.Next
		if copyNode.Next != nil {
			// 让复制节点的next指向复制节点的next的next，也就是原节点的next
			copyNode.Next = copyNode.Next.Next
		}
	}
	return copyHead
}

func TestRandNode(t *testing.T) {
	node1 := &RandomListNode{Val: 7}
	node2 := &RandomListNode{Val: 13}
	node3 := &RandomListNode{Val: 11}
	node4 := &RandomListNode{Val: 10}
	node5 := &RandomListNode{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1
	result := RandomListCopy(node1)
	for result != nil {
		t.Log(result.Val)
		if result.Random != nil {
			t.Log(result.Random.Val)
		}
		result = result.Next
	}
}
