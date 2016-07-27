package main

const host = "localhost:9001"

func main() {
	hub := NewHub(host)
	hub.Bind()
	hub.Listen()
}
