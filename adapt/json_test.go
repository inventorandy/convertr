package adapt

import "testing"

// TestUsingJSON is a test function for UsingJSON.
func TestUsingJSON(t *testing.T) {
	// Test Structures
	type StructA struct {
		FieldA int    `json:"field_a"`
		FieldB string `json:"field_b"`
	}
	type StructB struct {
		FieldA int    `json:"field_a"`
		FieldB string `json:"field_b"`
	}

	// Test Data
	a := StructA{FieldA: 42, FieldB: "Hello, World!"}
	var b StructB

	// Test UsingJSON
	if err := UsingJSON(a, &b); err != nil {
		t.Errorf("UsingJSON(a, &b) = %v; want nil", err)
	}
	if b.FieldA != 42 {
		t.Errorf("b.FieldA = %d; want 42", b.FieldA)
	}
	if b.FieldB != "Hello, World!" {
		t.Errorf("b.FieldB = %q; want %q", b.FieldB, "Hello, World!")
	}
}
