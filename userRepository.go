package main

import (
	"net"
	"sync"
)

type UserRepository struct {
	Users map[net.Conn]*User
	Mutex sync.RWMutex
}

func NewUserRepository() (*UserRepository) {
	repo := new(UserRepository)
	repo.Users = make(map[net.Conn]*User)

	return repo
}

func (repo *UserRepository) Add(conn net.Conn) {
	repo.Mutex.Lock()
	id := uint64(len(repo.Users) + 1)
	user := &User{ id, conn }
	repo.Users[conn] = user
	repo.Mutex.Unlock()
}

func (repo *UserRepository) GetByConn(conn net.Conn) (*User) {
	repo.Mutex.RLock()
	user := repo.Users[conn]
	repo.Mutex.RUnlock()

	return user
}

func (repo *UserRepository) GetAllByConnExcept(conn net.Conn) ([]*User) {
	repo.Mutex.RLock()	
	onlyUser := len(repo.Users) == 1
	
	if (onlyUser) {
		return make([]*User, 0)
	}

	users := make([]*User, len(repo.Users) - 1)
	i := 0

	for _, user := range repo.Users {
		if user.Conn != conn {
			users[i] = user
			i++
		}
	}

	repo.Mutex.RUnlock()	

	return users
}

func (repo *UserRepository) DeleteByConn(conn net.Conn) {
	repo.Mutex.Lock()
	delete(repo.Users, conn)
	repo.Mutex.Unlock()
}