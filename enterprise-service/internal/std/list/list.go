package list

import (
	"enterprise-service/internal/std"
	"errors"
	"fmt"
)

type node[T any] struct {
	next  *node[T]
	value T
}

// list structure to represent single linked lsit
type list[T std.StdComparable[T]] struct {
	head       *node[T]
	tail       *node[T]
	deleteType map[bool]func(l *list[T], n T)
}

// New creates new instance of linked list and returns pointer to it
//
// Pre-cond: set std.StdComparable generic type
//
// New instance of list is created
func New[T std.StdComparable[T]]() *list[T] {
	return &list[T]{
		head: nil,
		tail: nil,
	}
}

// Add adds given elem at the tail
//
// Pre-cond: given element to add
//
// Post-cond: list's tail now is equal to given element
func (l *list[T]) Add(item node[T]) {
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

// Find searches for given value in list
//
// Pre-cond: given value to find
//
// Post-cond: if value exists - finds it and returns value and nil error
// Otherwise returns default value and error
func (l *list[T]) Find(n T) (node[T], error) {
	var def T
	if l == nil || l.head == nil {
		return node[T]{}, fmt.Errorf("not found")
	}
	res := &node[T]{value: def, next: nil}

	tmp := l.head
	found := false

	pred := func() bool { return tmp != nil && !found }
	act := func() {
		if tmp.value.Compare(n) == 0 {
			res = tmp
			found = true
		}
		tmp = tmp.next
	}

	iter(pred, act)

	if !found {
		return *res, errors.New("not found")
	}
	return *res, nil
}

func iter(cond func() bool, action func()) {
	for cond() {
		action()
	}
}
