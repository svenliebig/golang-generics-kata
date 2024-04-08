package list

import "iter"

// Ideas are just commented out for now.

type list[T any] interface {
	// Add appends the element to the list.
	Add(i T) list[T]
	// Set sets the element at the given index.
	Set(i int, v T) list[T]
	// Get returns the element at the given index.
	Get(i int) T

	// filters the list based on the predicate p and returns a new list of elements that match the predicate.
	Filter(p func(item T) bool) list[T]
	// skips the first n elements and returns a list of the remaining elements.
	Skip(n int) list[T]
	// takes the first n elements and returns a list of them.
	Take(n int) list[T]

	// Map applies the function m to each element in the list and returns a new list of the results.
	Map(m func(item T) T) list[T]
	// Reduce(f func(acc T, item T) T) T
	// Reverse() list[T]
	// Sort() list[T]

	// IsSafe returns true if the list is safe to use concurrently.
	IsSafe() bool
	// ToSafe returns a safe list, that can be used concurrently.
	ToSafe() list[T]
	// ToUnsafe returns an unsafe list, that should not be used concurrently.
	ToUnsafe() list[T]

	// Iterator returns an iterator for the list elements.
	Iterator() iter.Seq[T]
	// Returns a slice of the list elements.
	Collect() []T
}
