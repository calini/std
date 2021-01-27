// WIP
package collection

type Element interface{}

type Iterator interface{}

type Collection interface {
	// ensures that this collection contains the specified element (optional operation)
	Add(e Element) bool
	// adds all of the elements in the specified collection to this collection (optional operation)
	AddAll(c Collection) bool
	// removes all of the elements from this collection (optional operation)
	Clear()
	// returns true if this collection contains the specified element
	Contains(e Element) bool
	// returns true if this collection contains all of the elements in the specified collection
	ContainsAll(c Collection) bool
	// returns true if this collection contains no elements
	IsEmpty() bool
	// returns an iterator over the elements contained in this collection
	Iterator() Iterator
	// removes a single instance of the specified element from this collection, if it is present (optional operation)
	Remove(e Element) bool
	// removes all of this collection's elements that are also contained in the specified collection (optional operation)
	RemoveAll(c Collection) bool
	// retains only the elements in this collection that are contained in the specified collection (optional operation)
	RetainAll(c Collection) bool
	// returns the number of elements in this collection
	Size() int
	// returns an array containing all of the elements in this collection
	ToArray() []Element
	// returns a string representation of this collection
	String() string
}