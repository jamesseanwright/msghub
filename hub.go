package main

import (
	"net"
	"encoding/json"
	"log"
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
	log.Println("Listening on", hub.Port)
}

func (hub *Hub) Listen() {
	defer hub.Listener.Close()

	for {
		conn, err := hub.Listener.Accept()
		log.Println("New connection!")

		if err != nil {
			panic(err)
		}

		hub.Actions.Users.Add(conn)
		go hub.ListenForRequests(conn)
	}
}

func (hub *Hub) ListenForRequests(conn net.Conn) {
	decoder := json.NewDecoder(conn);

	for decoder.More() {
		var request Request	
		decoder.Decode(&request)

		switch request.Type {
			case "getUserId":
				hub.Actions.GetUserId(conn)

			case "getAllUsers":
				hub.Actions.GetAllUsers(conn)

			case "logout":
				hub.Actions.Logout(conn)

			case "sendMessage":
				hub.Actions.SendMessage(conn, &request)

			default:
				hub.Actions.NotFound(conn)
		}
	}
}