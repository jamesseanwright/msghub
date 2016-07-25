package tests

import (
	"msghub"
	"testing"
	"net"
)

type MockConn struct {
	net.Conn
	CalledMethods map[string]bool
}

func NewMockConn() (*MockConn) {
	conn := new(MockConn)
	conn.CalledMethods = make(map[string]bool)
	return conn
}

func (conn *MockConn) Close() (error) {
	conn.CalledMethods["Close"] = true
	return nil
}

func (conn *MockConn) WasMethodCalled(name string) (bool) {
	return conn.CalledMethods[name]
}

func TestUserRepositoryAddAndGet(t *testing.T) {
	id := uint64(1)
	userRepository := main.NewUserRepository()
	userRepository.Add(NewMockConn());
	user := userRepository.Get(id)

	if user == nil {
		t.Fatal("Retrieved user is nil")
	}

	if user.Id != id {
		t.Fatalf("Unexpected user. Got %d, wanted %d", user.Id, id)
	}
}

func TestUserRepositoryGetAllExcept(t *testing.T) {
	id := uint64(2)
	userRepository := main.NewUserRepository()
	userRepository.Add(NewMockConn());
	userRepository.Add(NewMockConn());
	userRepository.Add(NewMockConn());
	
	users := userRepository.GetAllExcept(id)

	if users == nil {
		t.Fatal("Returned users map is nil")
	}

	if len(users) != 2 {
		t.Fatalf("Map is of incorrect length. Got %d, wanted %d", len(users), 2)
	}

	for _, user := range users {
		if user.Id == id {
			t.Fatal("Returned user for the ID that should have been excluded")			
		}
	}
}

func TestUserRepositoryDelete(t *testing.T) {
	id := uint64(1)
	userRepository := main.NewUserRepository()
	conn := NewMockConn()
	userRepository.Add(conn);
	userRepository.Delete(id);
	user := userRepository.Get(id)

	if user != nil {
		t.Fatal("User was not deleted")
	}

	if !conn.WasMethodCalled("Close") {
		t.Fatal("Close method was not called on user's connection")		
	}
}