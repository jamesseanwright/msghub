package main

import (
	"net"
	"encoding/json"
	"fmt"
)

type Hub struct {
	Port string
	Listener net.Listener
	Users *UserRepository
}

func NewHub(port string) (*Hub) {
	hub := new(Hub)
	hub.Port = port
	hub.Users = NewUserRepository()

	return hub
}

func (hub *Hub) Bind() {
	listener, err := net.Listen("tcp", hub.Port)

	if err != nil {
		panic(err)
	}

	hub.Listener = listener
	fmt.Println("Listening on", hub.Port)
}

func (hub *Hub) Listen() {
	for {
		conn, err := hub.Listener.Accept()

		if err != nil {
			panic(err)
		}

		hub.Users.Add(conn)

		go hub.ListenForRequests(conn)
	}
}

func (hub *Hub) ListenForRequests(conn net.Conn) {
	var request Request
	decoder := json.NewDecoder(conn);
	err := decoder.Decode(&request)

	if err != nil {
		panic(err)
	}

	/* Satisfying first functional test for now.
	 * Will split into router for next test */

	user := hub.Users.GetByConn(conn)
	encoder := json.NewEncoder(conn);	
	err = encoder.Encode(user)

	if err != nil {
		panic(err)
	}
}