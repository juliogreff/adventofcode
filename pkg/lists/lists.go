package lists

import "sort"

type sortable[T any] struct {
	list   []T
	lessFn func(T, T) bool
}

func (s sortable[T]) Len() int {
	return len(s.list)
}

func (s sortable[T]) Swap(i, j int) {
	a := s.list[i]
	b := s.list[j]

	s.list[i] = b
	s.list[j] = a
}

func (s sortable[T]) Less(i, j int) bool {
	return s.lessFn(s.list[i], s.list[j])
}

func Sort[T any](list []T, lessFn func(T, T) bool) {
	s := sortable[T]{
		list:   list,
		lessFn: lessFn,
	}

	sort.Sort(s)
}

func Reverse[T any](original []T) []T {
	reversed := make([]T, len(original))
	copy(reversed, original)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		tmp := len(reversed) - 1 - i
		reversed[i], reversed[tmp] = reversed[tmp], reversed[i]
	}

	return reversed
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	cp := make(map[K]V, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func FilterMapInPlace[K comparable, V any](m map[K]V, fn func(K, V) bool) {
	for k, v := range m {
		if !fn(k, v) {
			delete(m, k)
		}
	}
}
