package main

import (
	"encoding/json"
	"log"
	"net"
)

// Hub listens incoming TCP connections and determines
// which action should handle them
type Hub struct {
	Host     string
	Actions  *Actions
	Listener net.Listener
}

func NewHub(host string) *Hub {
	hub := new(Hub)
	hub.Host = host
	hub.Actions = NewActions()

	return hub
}

func (hub *Hub) Bind() {
	listener, err := net.Listen("tcp", hub.Host)

	if err != nil {
		panic(err)
	}

	hub.Listener = listener
	log.Println("Listening on", hub.Host)
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
		go hub.listenForRequests(conn)
	}
}

func (hub *Hub) listenForRequests(conn net.Conn) {
	decoder := json.NewDecoder(conn)

	for decoder.More() {
		var request Request
		decoder.Decode(&request)

		switch request.Type {
		case "getUserID":
			hub.Actions.GetUserID(conn)

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
