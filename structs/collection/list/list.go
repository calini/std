package list

import "github.com/calini/std/structs/collection"

type Element interface {
	collection.Element
}

type ListIterator interface{}

// Derived from java.util.AbstractList
type List interface {
	// appends the specified element to the end of this list (optional operation)
	Add(e Element) bool
	// inserts the specified element at the specified position in this list (optional operation)
	AddAtIndex(index int, e Element)
	// inserts all of the elements in the specified collection into this list at the specified position (optional operation)
	AddAll(index int, c ...Element) bool
	// removes all of the elements from this list (optional operation)
	Clear()
	// returns true if this list contains the specified element
	Contains(e Element)
	// returns true if this list contains all of the elements of the specified collection
	ContainsAll(c collection.Collection) bool
	// compares the specified object with this list for equality
	Equals(o interface{}) bool
	// returns the element at the specified position in this list
	Get(index int)
	// returns the hash code value for this list
	HashCode() int
	// returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element
	IndexOf(o Element) int
	// returns true if this list contains no elements
	IsEmpty() bool
	//// returns an iterator over the elements in this list in proper sequence
	//Iterator() collection.Iterator
	// returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element
	LastIndexOf(o Element) int
	//// returns a list iterator over the elements in this list (in proper sequence)
	//ListIterator() ListIterator
	//// returns a list iterator over the elements in this list (in proper sequence), starting at the specified position in the list
	//ListIteratorAtIndex(index int) ListIterator
	// removes the element at the specified position in this list (optional operation)
	RemoveIndex(index int) Element
	// removes the first occurrence of the specified element from this list, if it is present (optional operation)
	Remove(e Element) bool
	// removes from this list all of its elements that are contained in the specified collection (optional operation)
	RemoveAll(c collection.Collection) bool
	// retains only the elements in this list that are contained in the specified collection (optional operation)
	RetainAll(c collection.Collection) bool
	// removes from this list all of the elements whose index is between fromIndex, inclusive, and toIndex, exclusive
	RemoveRange(from, to int)
	// replaces the element at the specified position in this list with the specified element (optional operation)
	Set(index int, e Element) Element
	// returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive
	SubList(from, to int) List
	// returns an array containing all of the elements in this list in proper sequence (from first to last element)
	ToArray() []Element
	// returns a string representation of this collection
	String() string
}
