package client

import (
	"testing"
)

func TestClientFork(t *testing.T) {
	// Create parent
	parent := NewClient(ClientOptions{})

	// Set value in parent
	parent.Header.Set("a", "b")

	// Create child
	child := parent.Fork(ClientOptions{})

	if child.Header.Get("a") != "b" {
		t.Errorf("Child should inherit headers from parent")
	}

	child.Header.Set("c", "d")
	if parent.Header.Get("c") == "d" {
		t.Errorf("Parent should share not child's headers")
	}

	if !(child.Header.Get("a") == "b" && child.Header.Get("c") == "d") {
		t.Errorf("Child should be able to keep it's own headers")
	}
}
