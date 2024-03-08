package ptr

import (
	"reflect"
	"testing"
)

func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v, actual %#v", expected, actual)
	}
}

// TestFrom is a test function for From.
func TestFrom(t *testing.T) {
	equal(t, 42, *From(42))
}

// TestTo is a test function for To.
func TestTo(t *testing.T) {
	equal(t, 42, To(From(42)))
}

// TestFromSlice is a test function for FromSlice.
func TestFromSlice(t *testing.T) {
	equal(t, []*int{From(1), From(2), From(3)}, FromSlice([]int{1, 2, 3}))
}

// TestToSlice is a test function for ToSlice.
func TestToSlice(t *testing.T) {
	equal(t, []int{1, 2, 3}, ToSlice([]*int{From(1), From(2), From(3)}))
}
