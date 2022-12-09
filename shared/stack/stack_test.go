package stack

import "testing"

func TestStackSimple(t *testing.T) {
	stack := NewStack()

	stack.Push("X")
	stack.Push("Z")
	stack.Push("G")

	got := stack.Pop()
	want := "G"

	if want != got {
		t.Errorf("poped value: want %q, got %q", want, got)
	}

	if stack.length != 2 {
		t.Errorf("length: want %q, got %q", 2, stack.length)
	}
}

func TestPopFromEmpty(t *testing.T) {
	stack := NewStack()

	stack.Push("X")
	stack.Push("Z")
	stack.Push("G")

	stack.Pop()
	stack.Pop()
	stack.Pop()

	if stack.length != 0 {
		t.Errorf("Stack should be empty")
	}

	got := stack.Pop()
	want := ""

	if want != got {
		t.Errorf("Pop from empty: want %q, got %q", want, got)
	}
}
