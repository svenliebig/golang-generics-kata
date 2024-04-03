package list

// benchmarks with the map_list.go file show that similar performance profile
// as the native implementation with *pointer. :qa

type UntypedList struct {
	items []interface{}
}

func NewUntypedList() *UntypedList {
	return &UntypedList{
		items: make([]interface{}, 0),
	}
}

func (l *UntypedList) Add(item interface{}) {
	l.items = append(l.items, item)
}

func (l *UntypedList) Get(index int) interface{} {
	return l.items[index]
}

func (l *UntypedList) Map(m func(item interface{}) interface{}) *UntypedList {
	newList := NewUntypedList()

	for _, item := range l.items {
		newList.Add(m(item))
	}

	return newList
}
