package main

// Request defines the contract of a user's incoming request,
// which can contain target users and a message
type Request struct {
	Type    string
	UserIDs []uint64
	Message string
}
