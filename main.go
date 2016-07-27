package main

const port = ":9001"

func main() {
	hub := NewHub(port)
	hub.Bind()
	hub.Listen()
}
