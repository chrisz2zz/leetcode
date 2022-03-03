/**
* @Author: Chao
* @Date: 2022/3/3 10:15
* @Version: 1.0
 */

package main

import "fmt"

/*
todo 节点操作可以抽象为方法
 */
func main() {
	obj := Constructor(2)
	//fmt.Println(obj.Get(2))
	//obj.Put(2, 6)
	//fmt.Println(obj.Get(1))
	//obj.Put(1, 5)
	//obj.Put(1, 2)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(2))
	obj.Put(1, 0)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

type elem struct {
	key   int
	value int
	prev  *elem
	next  *elem
}

type LRUCache struct {
	capacity int
	used     int
	index    map[int]*elem
	head     *elem
	tail     *elem
}

func Constructor(capacity int) LRUCache {
	head := &elem{}
	tail := &elem{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		index:    make(map[int]*elem),
		head:     head,
		tail:     tail,
	}
}

/*
被访问的元素放到头部
*/
func (this *LRUCache) Get(key int) int {
	//判断元素是否存在
	elem, ok := this.index[key]
	if !ok {
		return -1
	}

	// 取出元素, 前驱指向后驱
	prev := elem.prev
	nex := elem.next
	prev.next = nex
	nex.prev = prev

	// 放到头部
	tmp := this.head.next
	elem.next = tmp
	this.head.next = elem
	tmp.prev = elem
	elem.prev = this.head

	return elem.value
}

/*
新元素放到链表头部
容量不足时  淘汰尾部
*/
func (this *LRUCache) Put(key int, value int) {
	// 特殊情况  若元素在链表中存在  删除原元素
	if v, ok := this.index[key]; ok {
		this.used--
		prev := v.prev
		nex := v.next
		prev.next = nex
		nex.prev = prev
		v = nil
		delete(this.index, key)
	}

	// 容量满的时候  移除尾部元素
	if this.used >= this.capacity  {
		//取出要淘汰的元素
		elim := this.tail.prev
		//淘汰元素的前驱(prev)
		prev := elim.prev
		//虚拟尾结点的前驱指向prev
		//prev.next指向虚拟尾结点
		//断开淘汰结点并置nil
		this.tail.prev = prev
		prev.next = this.tail
		delete(this.index, elim.key)
		elim = nil
		this.used--
	}

	this.used++

	e := &elem{
		key: key,
		value: value,
		prev:  nil,
		next:  nil,
	}

	nex := this.head.next
	nex.prev = e
	e.next = nex
	this.head.next = e
	e.prev = this.head
	this.index[key] = e
}
