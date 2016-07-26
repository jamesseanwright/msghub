package tests

import (
	"testing"
	"msghub"
)

func TestGetUserId(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.GetUserId(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestGetAllUsers(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.GetAllUsers(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestNotFound(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.NotFound(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}