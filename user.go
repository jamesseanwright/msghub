package main

import (
	"net"
)

// User represents an active TCP connection and its
// associated ID
type User struct {
	ID   uint64
	Conn net.Conn `json:"-"`
}
