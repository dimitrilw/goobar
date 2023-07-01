package set

type Void struct{}

type Set[T comparable] struct{ m map[T]Void }

func NewSet[T comparable](items []T) Set[T] {
	s := Set[T]{make(map[T]Void)}
	for _, item := range items {
		s.Add(item)
	}
	return s
}
func (s Set[T]) Len() int       { return len(s.m) }
func (s *Set[T]) Add(key T)     { (*s).m[key] = Void{} }
func (s Set[T]) Has(key T) bool { _, ok := s.m[key]; return ok }
func (s *Set[T]) Del(key T)     { delete((*s).m, key) }

/*
Built using an incrementing index (i) because it is faster than append.

Benchmark results:

	BenchmarkSet/SliceViaIncrement-10   31304653   37.05 ns/op   8 B/op   1 allocs/op
	BenchmarkSet/SliceViaAppend-10      28361586   42.09 ns/op   8 B/op   1 allocs/op
*/
func (s Set[T]) Slice() []T {
	res := make([]T, len(s.m))
	i := 0
	for k := range s.m {
		res[i] = k
		i++
	}
	return res
}
