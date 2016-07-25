package main

import (
	"net"
)

type User struct {
	Id uint64
	Conn net.Conn
}