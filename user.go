package main

import (
	"net"
)

type User struct {
	ID   uint64
	Conn net.Conn `json:"-"`
}
