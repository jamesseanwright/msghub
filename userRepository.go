package main

import (
	"net"
)

type UserRepository struct {
	Users map[net.Conn]*User
}

func NewUserRepository() (*UserRepository) {
	return &UserRepository{ make(map[net.Conn]*User) }
}

func (repo *UserRepository) Add(conn net.Conn) {
	id := uint64(len(repo.Users) + 1)
	user := &User{ id, conn }
	repo.Users[conn] = user
}

func (repo *UserRepository) GetByConn(conn net.Conn) (*User) {
	return repo.Users[conn]
}

func (repo *UserRepository) GetAllByConnExcept(conn net.Conn) ([]*User) {
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

	return users
}

func (repo *UserRepository) DeleteByConn(conn net.Conn) {
	conn.Close()
	delete(repo.Users, conn)
}