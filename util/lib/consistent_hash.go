//一致性哈希算法 Consistent Hashing
//wiki https://juejin.im/post/5ae1476ef265da0b8d419ef2
package lib

import (
	"bytes"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

//为解决数据倾斜问题（数据分布不均匀），引入虚拟节点机制
//对每一个节点计算多个哈希
const VIRTUAL_NODE = 32

//哈希环
type HashRing []uint32

//实现sort包的基础类型接口，便于使用sort对hash环进行排序、查找操作
//求长度
func (hr HashRing) Len() int {
	return len(hr)
}

//比较大小
func (hr HashRing) Less(i, j int) bool {
	return hr[i] < hr[j]
}

//数据交换
func (hr HashRing) Swap(i, j int) {
	hr[i], hr[j] = hr[j], hr[i]
}

//定义节点
type Node struct {
	//节点索引
	Id int
	//节点信息
	Info string
	//节点权重 @TODO 只能增大范围
	Weight int
}

//创建新的节点
func NewNode(id int, ip_port string, weight int) *Node {
	return &Node{
		Id:     id,
		Info:   ip_port,
		Weight: weight,
	}
}

//一致性哈希结构
type Consistent struct {
	//记录节点对应的位置
	Nodes map[uint32]Node
	//每个结点循环的次数
	numReps int
	//节点是否添加
	Resources map[int]bool
	//hash环
	ring HashRing
	sync.RWMutex
}

//创建哈希结构
func NewConsistent() *Consistent {
	return &Consistent{
		Nodes:     make(map[uint32]Node),
		numReps:   VIRTUAL_NODE,
		Resources: make(map[int]bool),
		ring:      HashRing{},
	}
}

//添加一个节点
func (c *Consistent) Add(node *Node) bool {
	//添加写入锁状态
	c.Lock()
	defer c.Unlock()

	//是否已经存在的节点
	if _, ok := c.Resources[node.Id]; ok {
		return false
	}

	//将一个节点根据权重均匀分布到hash环上
	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		//将当前节点放在hash环上
		str := c.stringConnect(i, node)
		c.Nodes[c.stringHash(str)] = *(node)
	}
	//标记此节点已经添加
	c.Resources[node.Id] = true
	//hash环重新排序
	c.hashRingSort()
	return true
}

//字符串连接
func (c *Consistent) stringConnect(i int, node *Node) string {
	var buf bytes.Buffer
	buf.WriteString(node.Info)
	buf.WriteString("*")
	buf.WriteString(strconv.Itoa(node.Weight))
	buf.WriteString(strconv.Itoa(i))
	buf.WriteString(strconv.Itoa(node.Id))
	return buf.String()
}

//字符串哈希值计算
func (c *Consistent) stringHash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

//对hash环进行排序
func (c *Consistent) hashRingSort() {
	c.ring = HashRing{}
	for k := range c.Nodes {
		c.ring = append(c.ring, k)
	}
	sort.Sort(c.ring)
}

//根据key返回指定节点
func (c *Consistent) Get(key string) Node {
	//添加读入锁状态
	c.RLock()
	defer c.RUnlock()

	hash := c.stringHash(key)
	i := c.search(hash)
	return c.Nodes[c.ring[i]]
}

//根据key查找指定的节点
func (c *Consistent) search(hash uint32) int {
	//如果未找到满足条件的节点则i=len(r.ring)
	i := sort.Search(len(c.ring), func(i int) bool {
		return c.ring[i] >= hash
	})

	if i < len(c.ring) {
		if i == len(c.ring)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1
	}
}

//删除指定节点
func (c *Consistent) Remove(node *Node) {
	//添加写入锁状态
	c.Lock()
	defer c.Unlock()

	//判断节点是否存在
	if _, ok := c.Resources[node.Id]; !ok {
		return
	}
	//删除节点标志
	delete(c.Resources, node.Id)

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.stringConnect(i, node)
		delete(c.Nodes, c.stringHash(str))
	}
	c.hashRingSort()
}
