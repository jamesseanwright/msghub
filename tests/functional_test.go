package tests

import (
	"testing"
	"net"
	"msghub"
	"strings"
	"encoding/json"
	"os"
)

const port = ":9001"

func TestMain(m *testing.M) {
	hub := main.NewHub(port)
	hub.Bind()
	go hub.Listen()
	os.Exit(m.Run())
}

func TestIdentityMessage(t *testing.T) {
	payload := `{ "type": "getUserId" }`
	conns := [3]net.Conn{ dial(port, t), dial(port, t), dial(port, t) }
	var user *main.User

	for i, conn := range conns {
		wantedId := i + 1
		sendPayload(conn, payload, t)
		unmarshal(&user, conn, t)
		
		if user == nil {
			t.Error("Expected user to not be nil")
		}

		if user.Id != uint64(wantedId) {
			t.Errorf("Expected user ID to be %d, but got %d", wantedId, user.Id)
		}

		conn.Close()
	}
}

func dial(port string, t *testing.T) (net.Conn) {
	conn, err := net.Dial("tcp", port)

	if err != nil {
		t.Fatalf("Couldn't connect to %s: %s", port, err)
	}

	return conn
}

func sendPayload(conn net.Conn, payload string, t *testing.T) {
	reader := strings.NewReader(payload)
	_, err := reader.WriteTo(conn)

	if err != nil {
		t.Error("Couldn't write data to connection:", err)
	}
}

func unmarshal(target interface{}, conn net.Conn, t *testing.T) {
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&target)

	if err != nil {
		t.Error("Couldn't deserialise JSON:", err)
	}
}