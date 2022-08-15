package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%d ", n.value)
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
		l.length++
		return
	}

	iterator := l.head // we are defining iterator outside for loop because we want to access it later on
	for ; iterator.next != nil; iterator = iterator.next {
	} // this gives us the last element of the linked list that we are adding the element to

	iterator.next = newNode
	l.length++
}

func (l *linkedList) remove(value int) {
	if l.head.value == value { // if the element is first element, then just point the head to its next element
		l.head = l.head.next
		l.length--
		return
	}

	previous := new(node) // allocates memory to new node that will keep track of previous node
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			previous.next = iterator.next
			l.length--
			return
		}
		previous = iterator // this will maintain previous node
	}
}

func (l linkedList) get(value int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%v ", iterator))
	}

	return sb.String()
}

func main() {
	fmt.Println("Welcome to the implementation of linked list in GoLang")

	list := linkedList{}
	list.add(32)
	list.add(43)
	list.add(24)
	list.add(69)

	fmt.Println(list)
	fmt.Println("Length of the linked list is", list.length)

	list.remove(43)
	fmt.Println(list)
	fmt.Println("Length of the linked list is", list.length)

	fmt.Println(list.get(69))

	list.add(-100)
	fmt.Println(list)
	fmt.Println("Length of the linked list is", list.length)

	list.remove(32)
	fmt.Println(list)
	fmt.Println("Length of the linked list is", list.length)
}
