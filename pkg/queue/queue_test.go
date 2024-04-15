package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()

	// Test IsEmpty
	if !queue.IsEmpty() {
		t.Error("Expected queue to be empty")
	}

	// Test Push and Length
	queue.Push(1)
	if queue.Length() != 1 {
		t.Error("Expected queue length to be 1")
	}

	// Test Peek
	if queue.Peek() != 1 {
		t.Error("Expected Peek to return 1")
	}

	// Test PeekTail
	if queue.PeekTail() != 1 {
		t.Error("Expected PeekTail to return 1")
	}

	// Test PushHead
	queue.PushHead(2)
	if queue.Length() != 2 {
		t.Error("Expected queue length to be 2")
	}

	// Test Peek after PushHead
	if queue.Peek() != 2 {
		t.Error("Expected Peek to return 2")
	}

	// Test PeekTail after PushHead
	if queue.PeekTail() != 1 {
		t.Error("Expected PeekTail to return 1")
	}

	// Test Pop
	if queue.Pop() != 2 {
		t.Error("Expected Pop to return 2")
	}

	// Test Length after Pop
	if queue.Length() != 1 {
		t.Error("Expected queue length to be 1")
	}

	// Test IsEmpty after Pop
	if queue.IsEmpty() {
		t.Error("Expected queue to not be empty")
	}
}
