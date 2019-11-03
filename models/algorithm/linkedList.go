//链表数据结构常见算法
package algorithm

import "fmt"

//@link https://www.bbsmax.com/A/Vx5M6LGvdN/
/*=====单链表翻转(start)=====*/
type ListNode struct {
	data interface{}
	Next *ListNode
}

//创建单向链表
func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.data = i
		cur = cur.Next
	}
}

func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}

func ReverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

/*=====单链表翻转(end)=====*/
