package list

import (
	"sync"
)

type TypedList[T interface{}] struct {
	items []T
	lock  sync.RWMutex
}

// func FromReader[T any](r io.Reader) *TypedList[T] {
// 	l := NewTypedList[T]()
// 	for {
// 		var bytes = make([]byte, 8)
// 		_, err := r.Read(bytes)
// 		if err == io.EOF {
// 			break
// 		}
// 		l.Add(bytes)
// 	}
// 	return l
// }

func From[T any](items []T) *TypedList[T] {
	return &TypedList[T]{
		items: items,
	}
}

func NewTypedList[T any]() *TypedList[T] {
	return &TypedList[T]{
		items: make([]T, 0),
	}
}

func NewTypedListWithCap[T any](c int) *TypedList[T] {
	return &TypedList[T]{
		items: make([]T, 0, c),
	}
}

func NewTypedListWithLen[T any](l int) *TypedList[T] {
	return &TypedList[T]{
		items: make([]T, l),
	}
}

func (l *TypedList[T]) Add(item T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.items = append(l.items, item)
}

func (l *TypedList[T]) Set(i int, item T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.items[i] = item
}

func (l *TypedList[T]) Get(index int) T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.items[index]
}

func (l *TypedList[T]) Filter(p func(item T) bool) *TypedList[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	// we could maybe use atleast capacity here
	newList := NewTypedList[T]()
	for _, item := range l.items {
		if p(item) {
			newList.Add(item)
		}
	}
	return newList
}

// This does not work, unfortunately

// func (l *TypedList[T]) Map[R any](m func(item T) R) *TypedList[R] {
// 	newList := NewTypedList[R]()
// 	for _, item := range l.items {
// 		newList.Add(m(item))
// 	}
// 	return newList
// }

func Map[T any, R any](l *TypedList[T], m func(item T) R) *TypedList[R] {
	l.lock.Lock()
	defer l.lock.Unlock()

	newList := NewTypedList[R]()
	for _, item := range l.items {
		newList.Add(m(item))
	}
	return newList
}

func MapWithLen[T any, R any](l *TypedList[T], m func(item T) R) *TypedList[R] {
	l.lock.Lock()
	defer l.lock.Unlock()

	newList := NewTypedListWithLen[R](len(l.items))
	for i, item := range l.items {
		newList.Set(i, m(item))
	}
	return newList
}

func MapWithCap[T any, R any](l *TypedList[T], m func(item T) R) *TypedList[R] {
	l.lock.Lock()
	defer l.lock.Unlock()

	newList := NewTypedListWithCap[R](len(l.items))
	for _, item := range l.items {
		newList.Add(m(item))
	}
	return newList
}
