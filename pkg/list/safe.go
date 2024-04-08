package list

import (
	"iter"
	"sync"
)

var (
	_ list[any] = &safe[any]{}
)

type safe[T any] struct {
	items []T
	lock  sync.RWMutex
}

func NewSafe[T any]() list[T] {
	return &safe[T]{}
}

func (l *safe[T]) Add(i T) list[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.items = append(l.items, i)

	return l
}

func (l *safe[T]) Set(i int, v T) list[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.items[i] = v

	return l
}

func (l *safe[T]) Get(i int) T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.items[i]
}

func (l *safe[T]) Filter(p func(item T) bool) list[T] {
	l.lock.RLock()
	defer l.lock.RUnlock()

	n := &safe[T]{
		items: make([]T, 0, len(l.items)),
	}

	for _, item := range l.items {
		if p(item) {
			n.items = append(n.items, item)
		}
	}

	return n
}

func (l *safe[T]) Map(m func(item T) T) list[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	n := &safe[T]{
		items: make([]T, 0, len(l.items)),
	}

	for _, item := range l.items {
		n.items = append(n.items, m(item))
	}

	return n
}

func (l *safe[T]) Skip(s int) list[T] {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if (s < 0) || (s >= len(l.items)) {
		return &safe[T]{
			items: make([]T, 0, len(l.items)),
		}
	}

	n := &safe[T]{
		items: make([]T, 0, len(l.items)-s),
	}

	for i := s; i < len(l.items); i++ {
		n.items = append(n.items, l.items[i])
	}

	return n
}

func (l *safe[T]) Take(t int) list[T] {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if (t < 0) || (t >= len(l.items)) {
		return &safe[T]{
			items: make([]T, 0, len(l.items)),
		}
	}

	n := &safe[T]{
		items: make([]T, 0, t),
	}

	for i := 0; i < t; i++ {
		n.items = append(n.items, l.items[i])
	}

	return n
}

func (l *safe[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {

		// I get a headache, while thinking about where to put this here, and when defer is called.

		// l.lock.RLock()
		// defer l.lock.RUnlock()

		for _, e := range l.items {
			if !yield(e) {
				return
			}
		}
	}
}

func (l *safe[T]) Collect() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.items
}

func (l *safe[T]) IsSafe() bool {
	return true
}

func (l *safe[T]) ToSafe() list[T] {
	return l
}

func (l *safe[T]) ToUnsafe() list[T] {
	return &unsafe[T]{
		items: l.items,
	}
}
