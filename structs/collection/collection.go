package collection

type Collection interface {
	// Ensures that this collection contains the specified element (optional operation).
	boolean	add(E e)
	// Adds all of the elements in the specified collection to this collection (optional operation).
	boolean	addAll(Collection<? extends E> c)
	// Removes all of the elements from this collection (optional operation).
	void	clear()
	// Returns true if this collection contains the specified element.
	boolean	contains(Object o)
	// Returns true if this collection contains all of the elements in the specified collection.
	boolean	containsAll(Collection<?> c)
	// Returns true if this collection contains no elements.
	boolean	isEmpty()
	// Returns an iterator over the elements contained in this collection.
	abstract Iterator<E>	iterator()
	// Removes a single instance of the specified element from this collection, if it is present (optional operation).
	boolean	remove(Object o)
	// Removes all of this collection's elements that are also contained in the specified collection (optional operation).
	boolean	removeAll(Collection<?> c)
	// Retains only the elements in this collection that are contained in the specified collection (optional operation).
	boolean	retainAll(Collection<?> c)
	// Returns the number of elements in this collection.
	abstract int	size()
	// Returns an array containing all of the elements in this collection.
	Object[]	toArray()
	// Returns an array containing all of the elements in this collection; the runtime type of the returned array is that of the specified array.
	<T> T[]	toArray(T[] a)
	// Returns a string representation of this collection.
	String	toString()
}