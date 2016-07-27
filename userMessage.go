package main

// UserMessage is a contract to present the
// content and sender of a relay message
type UserMessage struct {
	Message string
	From    uint64
}
