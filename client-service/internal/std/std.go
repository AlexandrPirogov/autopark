// std hold interfaces for all data structures that used in this project

package std

// Iterator incapsulate movement inside data structure elements in one way
type Iterator[T any] interface {
	// Next tells is there are elements to iterate
	//
	// if there are elements to iterate then return true
	// otherwise returnes false
	Next() bool

	// Curr returnes current element of iterator
	Curr() T
}

type Linked[T any] interface {
	// Add adds given elem at the tail
	//
	// Pre-cond: given element to add
	//
	// Post-cond: list's tail now is equal to given element
	PushBack(item T)

	// PopFront removes head of the list and returns item
	//
	// Pre-cond:
	//
	// Post-cond: if list isn't empty then returns value of head and true
	// and moves head to the next elem
	// Otherwise returns default value and false
	PopFront() (T, bool)

	// Iterator creates new instance of Iterator to inspect elements
	//
	// Pre-cond:
	//
	// Post-cond: returned instance of iterator
	Iterator() Iterator[T]
}

// AsSlice put all data structure's elemes into slice
//
// Pre-cond: given data structure's iterator
//
// Post-cond: returned slice of data structure's elements
func AsSlice[T any](i Iterator[T]) []T {
	res := make([]T, 0)
	for i.Next() {
		res = append(res, i.Curr())
	}
	return res
}
