//数据结构——链表
package datastructure

import (
	"fmt"
	"blog/util"
)

type Node struct {
	Data interface{}
	Next *Node
}

type MyLinkedList struct {
	Size int
	Head *Node //表头
	Tail *Node //表尾
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var list MyLinkedList = MyLinkedList{Size: 0, Head: nil, Tail: nil}
	return list
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= (*this).Size || index < 0 {
		return -1
	}

	item := (*this).Head
	for j := 0; j < index; j++ {
		item = (*item).Next
	}
	rs, err := util.Int((*item).Data)
	if err != nil {
		return -1
	}
	return rs
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := Node{Data: val, Next: nil}

	node.Next = (*this).Head
	(*this).Head = &node
	//当链表长度为0时
	if (*this).Size == 0 {
		(*this).Tail = &node
	}
	(*this).Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := Node{Data: val, Next: nil}

	tempNode := (*this).Tail
	(*tempNode).Next = &node
	//当链表长度为0时
	if (*this).Size == 0 {
		(*this).Head = &node
	}
	(*this).Tail = &node
	(*this).Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) bool {
	//index大于链表长度 则不进行插入
	if index > (*this).Size {
		return false
	}

	//如果小于0，则插入表头
	if index < 0 {
		index = 0
	}

	node := Node{Data: val, Next: nil}

	//插入表头
	if index == 0 {
		node.Next = (*this).Head
		(*this).Head = &node
		//当前链表长度为0
		if (*this).Size == 0 {
			(*this).Tail = &node
		}
	} else {
		var j int
		//找到i元素的前一个元素
		tempNode := (*this).Head
		for j = 1; j < index; j++ {
			tempNode = (*tempNode).Next
		}
		//原来元素放到新元素后面
		node.Next = (*tempNode).Next
		//新元素放到前一个元素后面
		(*tempNode).Next = &node

		//index为链表长度
		if index == (*this).Size {
			(*this).Tail = &node
		}
	}
	(*this).Size++
	return true
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) bool {
	if index >= (*this).Size || index < 0 {
		return false
	}

	//删除头部
	if index == 0 {
		node := (*this).Head
		(*this).Head = (*node).Next

		//如果只有一个元素，尾部需要一起调整
		if (*this).Size == 1 {
			(*this).Tail = nil
		}
	} else {
		//查找i节点前一个节点
		var j int
		tempNode := (*this).Head
		for j = 1; j < index; j++ {
			tempNode = (*tempNode).Next
		}

		//node := tempNode
		node := (*tempNode).Next
		(*tempNode).Next = (*node).Next

		//如果删除尾部元素
		if index == ((*this).Size - 1) {
			(*this).Tail = tempNode
		}
	}

	(*this).Size--
	return true
}

//打印整个链表
func (list *MyLinkedList) Print() {
	listLen := (*list).Size
	if listLen == 0 {
		fmt.Printf("当前链表为空!\n")
	}

	fmt.Printf("链表长度=%d\n", listLen)

	item := (*list).Head

	var j int
	for j = 0; j < listLen; j++ {
		fmt.Println(item.Data)
		item = (*item).Next
	}
}
