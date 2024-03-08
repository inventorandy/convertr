package ptr

// From returns a pointer to the given value.
func From[T any](t T) *T {
	return &t
}

// To returns the value pointed to by the given pointer.
func To[T any](t *T) T {
	return *t
}

// FromSlice iterates over the given slice and returns a slice of pointers to
// each element.
func FromSlice[T any](t []T) []*T {
	s := make([]*T, len(t))
	for i, v := range t {
		s[i] = &v
	}
	return s
}

// ToSlice iterates over the given slice of pointers and returns a slice of the
// values they point to.
func ToSlice[T any](t []*T) []T {
	s := make([]T, len(t))
	for i, v := range t {
		s[i] = *v
	}
	return s
}
