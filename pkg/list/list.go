package list

type list[T any] struct {
	items []T
}

func New[T any]() list[T] {
	return list[T]{}
}
