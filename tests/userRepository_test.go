package tests

import (
	"msghub"
	"testing"
)

func TestUserRepositoryAddAndGetByConn(t *testing.T) {
	id := uint64(1)
	userRepository := main.NewUserRepository()
	conn := NewMockConn()
	userRepository.Add(conn);
	user := userRepository.GetByConn(conn)

	if user == nil {
		t.Fatal("Retrieved user is nil")
	}

	if user.Id != id {
		t.Fatalf("Unexpected user. Got %d, wanted %d", user.Id, id)
	}
}

func TestUserRepositoryGetAllByConnExcept(t *testing.T) {
	id := uint64(2)
	conn := NewMockConn()
	userRepository := main.NewUserRepository()
	userRepository.Add(NewMockConn());
	userRepository.Add(conn);
	userRepository.Add(NewMockConn());
	
	users := userRepository.GetAllByConnExcept(conn)

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

func TestUserRepositoryDeleteByConn(t *testing.T) {
	userRepository := main.NewUserRepository()
	conn := NewMockConn()
	userRepository.Add(conn);
	userRepository.DeleteByConn(conn);
	user := userRepository.GetByConn(conn)

	if user != nil {
		t.Fatal("User was not deleted")
	}

	if !conn.WasMethodCalled("Close") {
		t.Fatal("Close method was not called on user's connection")		
	}
}