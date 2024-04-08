package list

import "iter"

var (
	_ list[any] = &unsafe[any]{}
)

type unsafe[T any] struct {
	items []T
}

func New[T any]() list[T] {
	return &unsafe[T]{}
}

func (l *unsafe[T]) Add(i T) list[T] {
	l.items = append(l.items, i)
	return l
}

func (l *unsafe[T]) Set(i int, v T) list[T] {
	l.items[i] = v
	return l
}

func (l *unsafe[T]) Get(i int) T {
	return l.items[i]
}

// Thoughts: Maybe I don't need a new list here, I could just work on the current list.
func (l *unsafe[T]) Filter(p func(item T) bool) list[T] {
	n := &unsafe[T]{
		items: make([]T, 0, len(l.items)),
	}

	for _, item := range l.items {
		if p(item) {
			n.Add(item)
		}
	}

	return n
}

// Thoughts: Maybe I don't need a new list here, I could just work on the current list.
func (l *unsafe[T]) Map(m func(item T) T) list[T] {
	n := &unsafe[T]{
		items: make([]T, 0, len(l.items)),
	}

	for _, item := range l.items {
		n.Add(m(item))
	}

	return n
}

// Thoughts: Maybe I don't need a new list here, I could just work on the current list.
func (l *unsafe[T]) Skip(s int) list[T] {
	if (s < 0) || (s >= len(l.items)) {
		return &unsafe[T]{
			items: make([]T, 0, len(l.items)),
		}
	}

	n := &unsafe[T]{
		items: make([]T, 0, len(l.items)-s),
	}

	for i := s; i < len(l.items); i++ {
		n.Add(l.items[i])
	}

	return n
}

// Thoughts: Maybe I don't need a new list here, I could just work on the current list.
func (l *unsafe[T]) Take(t int) list[T] {
	if (t < 0) || (t >= len(l.items)) {
		return &unsafe[T]{
			items: make([]T, 0, len(l.items)),
		}
	}

	n := &unsafe[T]{
		items: make([]T, 0, t),
	}

	for i := 0; i < t; i++ {
		n.Add(l.items[i])
	}

	return n
}

func (l *unsafe[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, e := range l.items {
			if !yield(e) {
				return
			}
		}
	}
}

func (l *unsafe[T]) Collect() []T {
	return l.items
}

func (l *unsafe[T]) IsSafe() bool {
	return false
}

func (l *unsafe[T]) ToSafe() list[T] {
	return &safe[T]{
		items: l.items,
	}
}

func (l *unsafe[T]) ToUnsafe() list[T] {
	return l
}
