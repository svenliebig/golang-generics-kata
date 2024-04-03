package list

import (
	"testing"
)

func BenchmarkOptimalSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]int, 1000)

		for i := 0; i < 1000; i++ {
			l[i] = i
		}

		l2 := make([]int, len(l))
		for i := 0; i < len(l); i++ {
			l2[i] = l[i] * 2
		}
	}
	b.ReportAllocs()
}

func BenchmarkOptimalTypedList(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := NewTypedListWithLen[int](1000)

		for i := 0; i < 1000; i++ {
			l.Set(i, i)
		}

		_ = MapWithLen(l, func(i int) int {
			return i * 2
		})
	}
	b.ReportAllocs()
}
