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

func (actions *Actions) GetAllUsers(conn net.Conn) (error) {
	user := actions.Users.GetAllByConnExcept(conn)
	encoder := json.NewEncoder(conn)
	return encoder.Encode(user)
}

func (actions *Actions) SendMessage(conn net.Conn, request *Request) (error) {
	sender := actions.Users.GetByConn(conn)
	recipients := actions.Users.GetByIds(request.UserIds)
	var err error

	for _, user := range recipients {
		if user == nil {
			message := ErrorMessage { "User(s) not found" }
			encoder := json.NewEncoder(conn)
			err = encoder.Encode(message)
			return err
		}

		userMessage := UserMessage { request.Message, sender.Id }
		encoder := json.NewEncoder(user.Conn)
		err = encoder.Encode(userMessage)
	}

	successMessage := InfoMessage { "Message delivered" }
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(successMessage)

	return err
}

func (actions *Actions) Logout(conn net.Conn) (error) {
	actions.Users.DeleteByConn(conn)
	info := InfoMessage { "Logged out" }
	encoder := json.NewEncoder(conn)
	return encoder.Encode(info)
}

func (actions *Actions) NotFound(conn net.Conn) (error) {
	err := ErrorMessage { "Command not found" }
	encoder := json.NewEncoder(conn)
	return encoder.Encode(err)
}