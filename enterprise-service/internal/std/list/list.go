package list

import "enterprise-service/std"

// listIterator type for linked list iterator
// that moves from head to tail
type listIterator[T any] struct {
	curr *node[T]
}

// new creates new instance of iterator for linked list.
// Take attention that value of iterator is pointer to node
func new[T any](val *node[T]) listIterator[T] {
	return listIterator[T]{
		curr: val,
	}
}

// Next tells is there are elements to iterate
//
// if there are elements to iterate then return true
// otherwise returnes false
func (l *listIterator[T]) Next() bool {
	return l.curr != nil
}

// Curr returnes current element of iterator
func (l *listIterator[T]) Curr() T {
	res := l.curr.value
	l.curr = l.curr.next
	return res
}

type node[T any] struct {
	next  *node[T]
	value T
}

// list structure to represent single linked lsit
type list[T any] struct {
	head       *node[T]
	tail       *node[T]
	deleteType map[bool]func(l *list[T], n T)
}

// New creates new instance of linked list and returns pointer to it
//
// Pre-cond: set std.StdComparable generic type
//
// New instance of list is created
func New[T any]() *list[T] {
	return &list[T]{
		head: nil,
		tail: nil,
	}
}

// Iterator creates new instance of Iterator to inspect elements
//
// Pre-cond:
//
// Post-cond: returned instance of iterator
func (l *list[T]) Iterator() std.Iterator[T] {
	var begin *node[T]
	if l == nil || l.head == nil {
		begin = nil
	} else {
		begin = l.head
	}
	return &listIterator[T]{curr: begin}
}

// Add adds given elem at the tail
//
// Pre-cond: given element to add
//
// Post-cond: list's tail now is equal to given element
func (l *list[T]) PushBack(t T) {
	item := node[T]{next: nil, value: t}
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

// Count returns number of elements
//
// Pre-cond:
//
// Post-cond: return number of elements in lsit
func (l *list[T]) Count() int {
	if l == nil {
		return 0
	}

	tmp := l.head
	count := 0

	pred := func() bool { return tmp != nil }
	act := func() {
		count++
		tmp = tmp.next
	}

	iter(pred, act)
	return count
}

// PopFront removes head of the list and returns item
//
// Pre-cond:
//
// Post-cond: if list isn't empty then returns value of head and true
// and moves head to the next elem
// Otherwise returns default value and false
func (l *list[T]) PopFront() (T, bool) {
	var def T
	if l == nil || l.head == nil {
		return def, false
	}

	def = l.head.value
	l.head = l.head.next
	return def, true
}

func iter(cond func() bool, action func()) {
	for cond() {
		action()
	}
}
