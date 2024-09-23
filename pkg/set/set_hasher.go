package set

type Hasher interface {
	Hash() string
}

type HasherSet[T Hasher] struct {
	m map[string]T
}

func NewHasherSet[T Hasher]() *HasherSet[T] {
	return &HasherSet[T]{m: map[string]T{}}
}

func (s *HasherSet[T]) Upsert(item T) {
	key := item.Hash()
	s.m[key] = item
}

func (s *HasherSet[T]) Get(item T) (out T, ok bool) {
	key := item.Hash()
	out, ok = s.m[key]
	return
}

func (s *HasherSet[T]) Iterator() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}
