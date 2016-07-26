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
		t.Fatal("Returned users array is nil")
	}

	if len(users) != 2 {
		t.Fatalf("array is of incorrect length. Got %d, wanted %d", len(users), 2)
	}

	for _, user := range users {
		if user.Id == id {
			t.Fatal("Returned user for the ID that should have been excluded")			
		}
	}
}

func TestUserRepositoryGetById(t *testing.T) {
	conn := NewMockConn()
	userRepository := main.NewUserRepository()
	userRepository.Add(conn);
	id := uint64(1)

	user := userRepository.GetById(id)

	if user == nil {
		t.Fatal("Returned user is nil")
	}

	if wanted, got := id, user.Id; wanted != got {
		t.Fatalf("Incorrect user returned. Wanted %d, got %d", wanted, got)
	}
}

func TestUserRepositoryGetByIdNilIfNotFound(t *testing.T) {
	userRepository := main.NewUserRepository()

	user := userRepository.GetById(1)

	if user != nil {
		t.Fatal("Expected returned user to be nil")
	}
}

func TestUserRepositoryGetByIds(t *testing.T) {
	userRepository := main.NewUserRepository()
	userRepository.Add(NewMockConn());
	userRepository.Add(NewMockConn());
	ids := []uint64{1, 2}

	users := userRepository.GetByIds(ids)

	if users == nil {
		t.Fatal("Returned users array is nil")
	}

	if wanted, got := len(ids), len(users); wanted != got {
		t.Fatalf("Returned array is wrong length. Wanted %d, got %d", wanted, got)
	}

	for _, user := range users {
		if user.Id != ids[0] && user.Id != ids[1] {
			t.Fatal("Returned incorrect user", user.Id)
		}
	}
}

func TestUserRepositoryOneUserGetAllByConnExcept(t *testing.T) {
	conn := NewMockConn()
	userRepository := main.NewUserRepository()
	userRepository.Add(conn);
	
	users := userRepository.GetAllByConnExcept(conn)

	if users == nil {
		t.Fatal("Returned users array is nil")
	}

	if len(users) != 0 {
		t.Fatalf("array is of incorrect length. Got %d, wanted %d", len(users), 0)
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
}