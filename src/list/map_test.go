package list

import (
	"testing"
)

type item struct {
	value int
}

func TestUntypedList(t *testing.T) {
	t.Run("create list", func(t *testing.T) {
		l := NewUntypedList()
		l.Add(item{value: 1})

		l2 := l.Map(func(i interface{}) interface{} {
			return item{value: i.(item).value * 2}
		})

		if len(l2.items) != 1 {
			t.Errorf("Expected list length to be 1, got %d", len(l2.items))
		}

		if l2.Get(0).(item).value != 2 {
			t.Errorf("Expected item value to be 2, got %d", l2.Get(0).(item).value)
		}
	})
}

func BenchmarkSliceForLen(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, item{value: i})
		}

		l2 := make([]item, 0)
		for i := 0; i < len(l); i++ {
			l2 = append(l2, item{value: l[i].value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForCapOnMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, item{value: i})
		}

		l2 := make([]item, 0, len(l))
		for i := 0; i < len(l); i++ {
			l2 = append(l2, item{value: l[i].value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForCapOnBoth(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]item, 0, 1000)

		for i := 0; i < 1000; i++ {
			l = append(l, item{value: i})
		}

		l2 := make([]item, 0, len(l))
		for i := 0; i < len(l); i++ {
			l2 = append(l2, item{value: l[i].value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, item{value: i})
		}

		l2 := make([]item, 0)
		for _, i := range l {
			l2 = append(l2, item{value: i.value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangeLenOnMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, item{value: i})
		}

		l2 := make([]item, len(l))
		for i, t := range l {
			l2[i] = item{value: t.value * 2}
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangeLenOnBoth(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]item, 1000)

		for i := 0; i < 1000; i++ {
			l[i] = item{value: i}
		}

		l2 := make([]item, len(l))
		for i, t := range l {
			l2[i] = item{value: t.value * 2}
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangeLenOnBothINT(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]int, 1000)

		for i := 0; i < 1000; i++ {
			l[i] = i
		}

		// I don't understand, when I remove this, until...
		l2 := make([]int, len(l))
		for i, t := range l {
			l2[i] = t * 2
		}

		if l2[500] != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500])
		}
		// ...here, I still only have 0 allocs/op. The first allocation of l is not considered, only the 2nd.
	}
	b.ReportAllocs()
}

func BenchmarkSliceForLenPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]*item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, &item{value: i})
		}

		l2 := make([]*item, 0)
		for i := 0; i < len(l); i++ {
			l2 = append(l2, &item{value: l[i].value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangePointer(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]*item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, &item{value: i})
		}

		l2 := make([]*item, 0)
		for _, i := range l {
			l2 = append(l2, &item{value: i.value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangePointerCapAtMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]*item, 0)

		for i := 0; i < 1000; i++ {
			l = append(l, &item{value: i})
		}

		l2 := make([]*item, 0, len(l))
		for _, i := range l {
			l2 = append(l2, &item{value: i.value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangePointerCapBoth(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]*item, 0, 1000)

		for i := 0; i < 1000; i++ {
			l = append(l, &item{value: i})
		}

		l2 := make([]*item, 0, len(l))
		for _, i := range l {
			l2 = append(l2, &item{value: i.value * 2})
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSliceForRangePointerLenOnBoth(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := make([]*item, 1000)

		for i := 0; i < 1000; i++ {
			l[i] = &item{value: i}
		}

		l2 := make([]*item, len(l))
		for i, t := range l {
			l2[i] = &item{value: t.value * 2}
		}

		if l2[500].value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2[500].value)
		}
	}
	b.ReportAllocs()
}

// Interesting is here, that the untyped list has similar performance as the
// native implementation with pointer.
func BenchmarkUntypedList(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := NewUntypedList()

		for i := 0; i < 1000; i++ {
			l.Add(item{value: i})
		}

		l2 := l.Map(func(i interface{}) interface{} {
			return item{value: i.(item).value * 2}
		})

		if l2.Get(500).(item).value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2.Get(500).(item).value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkTypedList(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := NewTypedList[item]()

		for i := 0; i < 1000; i++ {
			l.Add(item{value: i})
		}

		l2 := Map(l, func(i item) item {
			return item{value: i.value * 2}
		})

		if l2.Get(500).value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2.Get(500).value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkTypedListWithCapAtMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := NewTypedList[item]()

		for i := 0; i < 1000; i++ {
			l.Add(item{value: i})
		}

		l2 := MapWithCap(l, func(i item) item {
			return item{value: i.value * 2}
		})

		if l2.Get(500).value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2.Get(500).value)
		}
	}
	b.ReportAllocs()
}

func BenchmarkTypedListWithLenAtMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		l := NewTypedList[item]()

		for i := 0; i < 1000; i++ {
			l.Add(item{value: i})
		}

		l2 := MapWithLen(l, func(i item) item {
			return item{value: i.value * 2}
		})

		if l2.Get(500).value != 1000 {
			b.Errorf("Expected item value to be 1000, got %d", l2.Get(500).value)
		}
	}
	b.ReportAllocs()
}
