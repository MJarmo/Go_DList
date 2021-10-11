package main

import (
	"fmt"
)

type Node struct {
	prev, next *Node
	val int
}

type DList struct {
	first, last *Node
	size uint32
}

func (dl *DList) addToEmptyList(i int)  {
	dl.size++
	var node *Node
	node = &Node{prev: nil, next: nil, val: i}
	dl.first, dl.last = node, node
}

func (dl *DList) pushBack(i int)  {
	if dl.size == 0 {
		dl.addToEmptyList(i)
	} else {
		dl.size++
		var node *Node
		node = &Node{prev: dl.last, next: nil, val: i}
		dl.last.next = node
		dl.last = node
	}
}

func (dl *DList) pushFront(i int) {
	if dl.size == 0 {
		dl.addToEmptyList(i)
	} else {
		dl.size++
		var node *Node
		node = &Node{prev: nil, next: dl.first, val: i}
		dl.first.prev, dl.first = node, node

	}
}

func (dl *DList) insert(val int, place uint32) {
	if place > dl.size {
		fmt.Println("U did fuck up this")
		return
	}
	
	if place == 0 {
		dl.pushFront(val)
		return
	}
	
	if place == dl.size {
		dl.pushBack(val)
		return
	}
	
	if dl.size == 0 {
		dl.addToEmptyList(val)
	} else {
		var n *Node
		n = dl.first
		for i := uint32(0); i < dl.size; i++ {
			if i == place-1 {
				break
			}
			n = n.next
		}
		
		dl.size++
		var newNode *Node
		newNode = &Node{prev: n.prev, next: n, val: val}
		n.prev.next, newNode.prev = newNode , n.prev.next
		n.prev, newNode.next = newNode , n
		
	}
}

func (dl *DList) popBack() {
	dl.size--
	dl.last, dl.last.next = dl.last.prev, nil
}

func (dl *DList) popFront() {
	dl.size--
	dl.first, dl.first.prev = dl.first.next, nil
}

func (dl *DList) remove(i uint32) {
	if i> dl.size {
		fmt.Println("U did fuck up this")
		return
	}
	
	if i== 0 {
		dl.popFront()
		return
	}
	
	if i== dl.size {
		dl.popBack()
		return
	}
}

func (dl *DList) printList() {
	var next *Node
	next = dl.first
	for i := uint32(0); i<dl.size; i++ {
		fmt.Println(next)
		next = next.next
	}
	fmt.Println(dl.size)
}



func main() {

	
	var dl DList
	
	dl.pushBack(8)
	dl.pushBack(2)
	dl.pushBack(5)
	dl.pushFront(56)
	dl.pushFront(56)
	dl.insert(66, 0)
	dl.printList()
	fmt.Println("SDF")
	dl.popFront()
	dl.printList()

}
