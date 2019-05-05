package pie

import (
	"testing"
)

// Make sure that Append never alters the receiver, or other
// slices sharing the same memory, unlike the built-in append.
func TestAppendNonDestructive(t *testing.T) {
	ab := Strings{"A", "B"}
	if x, expected := ab.Join(""), "AB"; x != expected {
		t.Errorf("Expected %q, got %q", expected, x)
	}

	abc := ab.Append("C")
	aby := ab.Append("Y")
	if x, expected := abc.Join(""), "ABC"; x != expected {
		t.Errorf("Expected %q, got %q", expected, x)
	}
	if x, expected := aby.Join(""), "ABY"; x != expected {
		t.Errorf("Expected %q, got %q", expected, x)
	}

	abcd := abc.Append("D")
	abcz := abc.Append("Z")
	if x, expected := abcd.Join(""), "ABCD"; x != expected {
		t.Errorf("Expected %q, got %q", expected, x)
	}
	if x, expected := abcz.Join(""), "ABCZ"; x != expected {
		t.Errorf("Expected %q, got %q", expected, x)
	}
}
