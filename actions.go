package main

import (
	"net"
	"encoding/json"
)

type Actions struct {
	Users *UserRepository
}

func NewActions() (*Actions) {
	actions := new(Actions)
	actions.Users = NewUserRepository()
	return actions
}

func (actions *Actions) GetUserId(conn net.Conn) (error) {
	user := actions.Users.GetByConn(conn)
	encoder := json.NewEncoder(conn)
	return encoder.Encode(user)
}

func (actions *Actions) NotFound(conn net.Conn) (error) {
	err := ErrorMessage { "Command not found" }
	encoder := json.NewEncoder(conn)
	return encoder.Encode(err)
}