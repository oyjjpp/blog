//数据结构——链表
package datastructure

import "fmt"

//单节点结构
type Node struct {
	data interface{}
	Next *Node
}

//链表结构
type List struct {
	size uint64 //链表数量
	head *Node  //表头
	tail *Node  //表尾
}

//链表初始化操作
func (list *List) Init() {
	(*list).size = 0   //空链表
	(*list).head = nil //设置表头为nil
	(*list).tail = nil //设置表尾为nil
}

//顺序向链表尾部添加元素
//如果原链表长度为0 则需要设置表头
//如果原链表长度>0 则继续向表尾添加元素
func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	//设置添加的节点下一节点为空
	(*node).Next = nil

	//将新元素放入单链表中
	if (*list).size == 0 {
		(*list).head = node
	} else {
		tmpNode := (*list).tail
		(*tmpNode).Next = node
	}

	//调整尾部节点，及链表元素数量
	(*list).tail = node
	(*list).size++

	return true
}

//任意位置插入元素
//i==0 则调整链表 表头
//i>0 则需要通过循环找到i元素的前一个元素
func (list *List) Insert(i uint64, node *Node) bool {
	//空的节点、索引超出范围、空链表都无法做插入操作
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}

	//插入表头
	if i == 0 {
		(*node).Next = (*list).head
		(*list).head = node
	} else {
		var j uint64
		//找到i元素的前一个元素
		tempNode := (*list).head
		for j = 1; j < i; j++ {
			tempNode = (*tempNode).Next
		}
		//原来元素的后一个位置放到新元素后面
		(*node).Next = (*tempNode).Next
		//新元素放到前一个元素后面
		(*tempNode).Next = node
	}
	(*list).size++
	return true
}

//删除指定位置元素
//注意尾部节点的问题
func (list *List) Remove(i uint64) bool {
	if i >= (*list).size {
		return false
	}

	//删除头部
	if i == 0 {
		node := (*list).head
		(*list).head = (*node).Next

		//如果只有一个元素，尾部需要一起调整
		if (*list).size == 1 {
			(*list).tail = nil
		}
	} else {
		//查找i节点前一个节点
		var j uint64
		tempNode := (*list).head
		for j = 1; j < i; j++ {
			tempNode = (*tempNode).Next
		}

		node := (*tempNode).Next
		(*tempNode).Next = (*node).Next

		//如果删除尾部元素
		if i == ((*list).size - 1) {
			(*list).tail = tempNode
		}
	}

	(*list).size--
	return true
}

//删除链表的所有元素
func (list *List) RemoveAll() bool {
	if (*list).size == 0 {
		return false
	}

	//防止内存泄漏，逐个循环删除元素
	for (*list).head != nil {
		(*list).head = (*list).head.Next
	}
	(*list).head = nil
	(*list).tail = nil
	return true
}

//获取某个位置的元素
func (list *List) Get(i uint64) *Node {
	//超过链表大小
	if i >= (*list).size {
		return nil
	}

	var j uint64
	item := (*list).head
	for j = 0; j < i; j++ {
		item = (*item).Next
	}
	return item
}

//获取链表长度
func (list *List) GetSize() uint64 {
	return (*list).size
}

//打印整个链表
func (list *List) Print() {
	listLen := (*list).size
	if listLen == 0 {
		fmt.Printf("当前链表为空!\n")
	}

	fmt.Printf("链表长度=%d\n", listLen)

	item := (*list).head

	var j uint64
	for j = 0; j < listLen; j++ {
		fmt.Println(item.data)
		item = (*item).Next
	}
}
