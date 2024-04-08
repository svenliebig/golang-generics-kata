package main

import (
	"fmt"

	"github.com/svenliebig/golang-generics-kata/pkg/list"
)

func main() {
	unsafe := list.New[int]()

	unsafe.Add(1)
	unsafe.Add(2)
	unsafe.Add(3)

	safe := unsafe.ToSafe()

	safe.Add(4)
	safe.Add(5)
	safe.Add(6)

	unsafe = safe.ToUnsafe()

	for v := range unsafe.Iterator() {
		fmt.Print(v)
	}

	fmt.Println()

	unsafe = unsafe.Filter(func(item int) bool {
		return item%2 == 0
	})

	for v := range unsafe.Iterator() {
		fmt.Print(v)
	}

	fmt.Println()

	unsafe = unsafe.Map(func(item int) int {
		return item * 2
	})

	for v := range unsafe.Iterator() {
		fmt.Print(v)
	}

	fmt.Println()

	unsafe = unsafe.Skip(1)

	for v := range unsafe.Iterator() {
		fmt.Print(v)
	}

	fmt.Println()

	unsafe = unsafe.Take(1)

	for v := range unsafe.Iterator() {
		fmt.Print(v)
	}

	fmt.Println()
}
