package main

import (
	"sync"
)

const port = ":9001"


func main() {
	var wg sync.WaitGroup
	hub := NewHub(port)
	hub.Bind()
	go hub.Listen()

	wg.Add(1)	
	wg.Wait()
}