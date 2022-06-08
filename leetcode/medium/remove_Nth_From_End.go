package main

import (
	"fmt"
)

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	fmt.Println()
	newHead := removeNthFromEndLast(&head, 1)
	curr := newHead
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 删除链表的倒数第 N 个结点 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
* @link https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	curr := *head
	for curr.Next != nil {
		curr = *curr.Next
		length++
	}
	if n > length {
		return head
	}
	if n == length {
		head = head.Next
		return head
	}

	count := 0
	pre := head
	newCurr := head
	for newCurr.Next != nil {
		if count == length-n {
			break
		}
		pre = newCurr
		newCurr = newCurr.Next
		count++
	}
	pre.Next = pre.Next.Next
	return head
}

func removeNthFromEndLast(head *ListNode, n int) *ListNode {
	// p-变量，指针变量指向的内存地址 &p-编译器为变量p分配的内存地址 *p-p指针指向的内容
	dummy := &ListNode{0, head}
	// 第一个首节点
	first := head
	// 第二个首节点
	second := dummy
	// 遍历第一个首节点n次 则剩余可遍历次数为len-n次
	for i := 0; i < n; i++ {
		first = first.Next
	}
	// 遍历首节点到最后一个 并且移动第二个首节点，当首节点为nil时则说明找到需要删除元素的前一个节点
	for first != nil {
		first = first.Next
		second = second.Next
	}
	// 删除元素
	second.Next = second.Next.Next
	// 返回节点
	return dummy.Next
}
