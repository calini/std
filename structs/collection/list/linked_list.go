package list

import (
	"fmt"
	"github.com/calini/std/structs/collection"
)

type Node struct {
	next *Node
	Val  interface{}
}

type LinkedList struct {
	first *Node
	last *Node
	size int
}

// Add appends the specified element to the end of this list
func (l *LinkedList) Add(e interface{}) {
	newNode := &Node{nil, e}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
	} else {
		l.last.next = newNode
		l.last = newNode
	}
	l.size++
}

// Next returns the next node in line, if exists, else errors out
func (n *Node) Next() (*Node, error) {
	if n == nil {
		return nil, fmt.Errorf("node nil")
	}

	if n.next == nil {
		return nil, fmt.Errorf("last node")
	}

	return n.next, nil
}

// Seek returns the node that is 'index' positions away, if possible, else errors out
func (n *Node) Seek(index int) (*Node, error) {
	ptr := n
	for i := 1; i < index; i++ {
		var err error
		ptr, err = ptr.Next()
		if err != nil {
			return nil, fmt.Errorf("index out of bounds")
		}
	}

	return ptr, nil
}

// AddAtIndex inserts the specified element at the specified position in this list
func (l *LinkedList) AddAtIndex(index int, e interface{}) error {
	newNode := &Node{nil, e}

	if index == 0 {
		newNode.next = l.first // may be null
		l.first = newNode
		if l.last == nil {
			l.last = newNode
		}
		return nil
	} else {
		ptr, err := l.first.Seek(index)
		if err != nil {
			return err
		}

		newNode.next = ptr.next
		ptr.next = newNode

		if index == l.size {
			l.last = newNode
		}

		l.size++
	}
	return nil
}

// AddAll adds a list of elements at the end of the list
func (l *LinkedList) AddAll(c ...interface{}) error {
	for _, elem := range c {
		l.Add(elem)
	}
	return nil
}

// AddAllAtIndex adds a list of elements at the specified index
func (l *LinkedList) AddAllAtIndex(index int, c ...Element) error {
	panic("implement me")
}
// AddFirst inserts the specified element at the beginning of this list
func (l *LinkedList) AddFirst(e Element) bool {
	panic("implement me")
}

// AddLast inserts the specified element at the end of this list
func (l *LinkedList) AddLast(e Element) bool {
	panic("implement me")
}

// Clear removes all of the elements from this list
func (l *LinkedList) Clear() {
	l.first, l.last = nil, nil
}

// Contains checks if the element e exists in the list
func (l *LinkedList) Contains(e interface{}) bool {
	ptr := l.first
	for ptr != nil {
		if ptr.Val == e {
			return true
		}
		ptr = ptr.next
	}

	return false
}

func (l *LinkedList) ContainsAll(c collection.Collection) bool {
	panic("implement me")
}

func (l *LinkedList) Equals(l2 *LinkedList) bool {
	panic("implement me")
}

func (l *LinkedList) Get(index int) (*Node, error) {
	return l.first.Seek(index)
}

func (l *LinkedList) HashCode() int {
	panic("implement me")
}

func (l *LinkedList) IndexOf(o Element) int {
	panic("implement me")
}

func (l *LinkedList) IsEmpty() bool {
	panic("implement me")
}

func (l *LinkedList) LastIndexOf(o Element) int {
	panic("implement me")
}

func (l *LinkedList) RemoveIndex(index int) Element {
	panic("implement me")
}

func (l *LinkedList) Remove(e Element) bool {
	panic("implement me")
}

func (l *LinkedList) RemoveAll(c collection.Collection) bool {
	panic("implement me")
}

func (l *LinkedList) RetainAll(c collection.Collection) bool {
	panic("implement me")
}

func (l *LinkedList) RemoveRange(from, to int) {
	panic("implement me")
}

func (l *LinkedList) Set(index int, e Element) Element {
	panic("implement me")
}

func (l *LinkedList) SubList(from, to int) List {
	panic("implement me")
}

func (l *LinkedList) ToArray() []Element {
	panic("implement me")
}

func (l *LinkedList) String() string {
	panic("implement me")
}
