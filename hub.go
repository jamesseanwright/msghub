package main

import (
	"net"
	"encoding/json"
	"fmt"
)

type Hub struct {
	Port string
	Actions *Actions
	Listener net.Listener
}

func NewHub(port string) (*Hub) {
	hub := new(Hub)
	hub.Port = port
	hub.Actions = NewActions()

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

		hub.Actions.Users.Add(conn)

		go hub.ListenForRequests(conn)
	}
}

func (hub *Hub) ListenForRequests(conn net.Conn) {
	var request Request
	decoder := json.NewDecoder(conn);

	for decoder.More() {
		err := decoder.Decode(&request)

		if err != nil {
			panic(err)
		}

		switch request.Type {
			case "getUserId":
				hub.Actions.GetUserId(conn)
				break

			case "getAllUsers":
				hub.Actions.GetAllUsers(conn)

			case "disconnect":
				hub.Actions.Disconnect(conn)

			default:
				hub.Actions.NotFound(conn)
		}
	}
}