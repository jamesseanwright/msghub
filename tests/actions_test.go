package tests

import (
	"testing"
	"msghub"
)

func TestGetUserId(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.Users.Add(conn) // TODO: mock user repo
	actions.GetUserId(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestGetAllUsers(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.Users.Add(conn)
	actions.Users.Add(NewMockConn())
	actions.GetAllUsers(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestDisconnect(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.Users.Add(conn)
	actions.Disconnect(conn)

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
