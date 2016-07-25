package main

import (
	"net"
)

type UserRepository struct {
	Users map[uint64]*User
}

func NewUserRepository() (*UserRepository) {
	return &UserRepository{ make(map[uint64]*User) }
}

func (repo *UserRepository) Add(conn net.Conn) {
	id := uint64(len(repo.Users) + 1)
	user := &User{ id, conn }
	repo.Users[id] = user
}

func (repo *UserRepository) Get(id uint64) (*User) {
	return repo.Users[id]
}

func (repo *UserRepository) GetAllExcept(id uint64) ([]*User) {
	users := make([]*User, len(repo.Users) - 1)
	i := 0

	for _, user := range repo.Users {
		if user.Id != id {
			users[i] = user
			i++
		}
	}

	return users
}

func (repo *UserRepository) Delete(id uint64) {
	repo.Users[id].Conn.Close()
	delete(repo.Users, id)
}