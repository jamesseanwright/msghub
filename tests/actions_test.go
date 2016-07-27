package tests

import (
	"testing"
	"msghub"
)

func TestActionsGetUserID(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.Users.Add(conn) // TODO: mock user repo
	actions.GetUserID(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestActionsGetAllUsers(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.Users.Add(conn)
	actions.Users.Add(NewMockConn())
	actions.GetAllUsers(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestActionsLogout(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.Users.Add(conn)
	actions.Logout(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestActionsSendMessage(t *testing.T) {
	conn := NewMockConn()
	request := &main.Request{ "sendMessage", []uint64{2}, "Hello" }
	actions := main.NewActions()
	actions.Users.Add(conn)
	actions.Users.Add(NewMockConn())
	actions.SendMessage(conn, request)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}

func TestActionsNotFound(t *testing.T) {
	conn := NewMockConn()
	actions := main.NewActions()
	actions.NotFound(conn)

	if !conn.WasMethodCalled("Write") {
		t.Fatal("Connection was never written to")
	}
}
