package main

import (
	"encoding/json"
	"net"
)

// Actions contains various handlers for responding to incoming TCP requests
type Actions struct {
	Users *UserRepository
}

func NewActions() *Actions {
	actions := new(Actions)
	actions.Users = NewUserRepository()
	return actions
}

func (actions *Actions) GetUserID(conn net.Conn) error {
	user := actions.Users.GetByConn(conn)
	encoder := json.NewEncoder(conn)
	return encoder.Encode(user)
}

func (actions *Actions) GetAllUsers(conn net.Conn) error {
	user := actions.Users.GetAllByConnExcept(conn)
	encoder := json.NewEncoder(conn)
	return encoder.Encode(user)
}

func (actions *Actions) SendMessage(conn net.Conn, request *Request) error {
	sender := actions.Users.GetByConn(conn)
	recipients := actions.Users.GetByIDs(request.UserIDs)
	var err error

	for _, user := range recipients {
		if user == nil {
			message := Message{"User(s) not found"}
			encoder := json.NewEncoder(conn)
			err = encoder.Encode(message)
			return err
		}

		userMessage := UserMessage{request.Message, sender.ID}
		encoder := json.NewEncoder(user.Conn)
		err = encoder.Encode(userMessage)
	}

	successMessage := Message{"Message delivered"}
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(successMessage)

	return err
}

func (actions *Actions) Logout(conn net.Conn) error {
	actions.Users.DeleteByConn(conn)
	info := Message{"Logged out"}
	encoder := json.NewEncoder(conn)
	return encoder.Encode(info)
}

func (actions *Actions) NotFound(conn net.Conn) error {
	err := Message{"Command not found"}
	encoder := json.NewEncoder(conn)
	return encoder.Encode(err)
}
