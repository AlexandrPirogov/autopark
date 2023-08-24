package std

type Linked[T StdComparable[T]] interface {
	Add(item T)
}

type StdComparable[T any] interface {
	comparable
	Compare(with T) int
}
