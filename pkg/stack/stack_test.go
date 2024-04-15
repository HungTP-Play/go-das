package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()

	// Test IsEmpty
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	// Test Push and Length
	stack.Push(1)
	if stack.Length() != 1 {
		t.Error("Expected stack length to be 1")
	}

	// Test Peek
	if stack.Peek() != 1 {
		t.Error("Expected Peek to return 1")
	}

	// Test PeekTail
	if stack.PeekTail() != 1 {
		t.Error("Expected PeekTail to return 1")
	}

	// Test PushHead
	stack.PushHead(2)
	if stack.Length() != 2 {
		t.Error("Expected stack length to be 2")
	}

	// Test Peek after PushHead
	if stack.Peek() != 1 {
		t.Error("Expected Peek to return 1")
	}

	// Test PeekTail after PushHead
	if stack.PeekTail() != 2 {
		t.Error("Expected PeekTail to return 2")
	}

	// Test Pop
	if stack.Pop() != 1 {
		t.Error("Expected Pop to return 1")
	}

	// Test Length after Pop
	if stack.Length() != 1 {
		t.Error("Expected stack length to be 1")
	}

	// Test IsEmpty after Pop
	if stack.IsEmpty() {
		t.Error("Expected stack to not be empty")
	}
}
