package questions

import (
	"github.com/xianyunyh/Learn-Algorithms/Go_codes/dataStruct"
)

//输入一个单链表，输出链表的倒数第K个节点

func FindReserveKNode(l *dataStruct.LinkedList, k int) *dataStruct.Node  {
	if l.Head == nil || k == 0 {
		return nil
	}
	slowNode := l.Head
	fastNode := l.Head
	for i := 1; i <= k ; i++  {
		if  fastNode != nil {
			fastNode = fastNode.Next
		}
	}
	for fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
	return slowNode
}

// 判断链表是否有环
/**
快慢指针 快的先走一步 然后慢的再走 如果有环 两个指针 一定在环内相遇 没环在终点相遇
 */

func LinkExistLoop(l *dataStruct.LinkedList) bool  {
	fast,slow := l.Head,l.Head
	for slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//节点相遇
		if slow == fast {
			return true
		}
	}
	return false
}