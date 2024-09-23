package set

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: map[T]struct{}{}}
}

func (s *Set[T]) Upsert(item T) {
	s.m[item] = struct{}{}
}

func (s *Set[T]) Get(item T) (out T, ok bool) {
	_, ok = s.m[item]
	if !ok {
		return
	}
	return item, true
}

func (s *Set[T]) Iterator() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for k := range s.m {
			if !yield(k) {
				return
			}
		}
	}
}
