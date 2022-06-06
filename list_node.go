package main

import (
	"fmt"
)

//单链表反转
//链表中环的检测
//两个有序的链表合并
//删除链表倒数第 n 个结点
//求链表的中间结点

type ListNode struct {
	Val  int
	Next *ListNode
}

func ShowNode(node *ListNode) {
	for node != nil {
		fmt.Println(*node)
		node = node.Next
	}
}

func Append(node *ListNode, ele int) {
	tail := *node
	var newNode = &ListNode{Val: ele}
	newNode.Next = &tail
	*node = *newNode
}

func Push(node *ListNode, ele int) {
	var newNode = &ListNode{Val: ele}
	for {
		if node.Next == nil {
			node.Next = newNode
			break
		}
		node = node.Next
	}
}

func Reverse(node *ListNode) {
	var head *ListNode
	head = &ListNode{
		Val:  node.Val,
		Next: nil,
	}
	var p *ListNode
	p = node.Next
	for true {
		if p == nil {
			break
		}
		Append(head, p.Val)
		p = p.Next
	}
	*node = *head
}

func Delete(node *ListNode, pos int) {
	length := Length(node)
	if pos > length-1 {
		return
	}
	index := 0
	var p *ListNode
	var head *ListNode
	p = node
	head = node
	for {
		if p == nil {
			break
		}
		if pos == 0 {
			*node = *node.Next
			break
		}
		if index == pos {
			if index == length-1 {
				head.Next = nil
				p = head.Next
				break
			}
			head.Next = p.Next
			p = head.Next
			break
		}
		p = p.Next
		if index > 0 {
			head = head.Next
		}
		index = index + 1
	}
}

func Length(node *ListNode) int {
	i := 0
	for {
		i = i + 1
		if node.Next == nil {
			return i
		}
		node = node.Next
	}
}

func main() {
	node := &ListNode{
		Val:  0,
		Next: nil,
	}
	Append(node, 1)
	Push(node, 2)
	Append(node, 3)
	Append(node, 4)
	Append(node, 5)
	fmt.Println("length", Length(node))
	ShowNode(node)
	//fmt.Println("delete----", 0)
	//Delete(node, 0)
	//ShowNode(node)
	//fmt.Println("delete----", "last")
	//Delete(node, Length(node))
	//ShowNode(node)
	fmt.Println("delete---", "last")
	Delete(node, Length(node)-1)
	ShowNode(node)

	fmt.Println("delete---", "head")
	Delete(node, 0)
	ShowNode(node)

	fmt.Println("delete---", "2")
	Delete(node, 2)
	ShowNode(node)

	fmt.Println("reverse----")
	Reverse(node)
	ShowNode(node)
}
