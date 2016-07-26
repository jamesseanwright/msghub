package tests

import (
	"encoding/json"
	"net"
	"os"
	"strings"
	"testing"
	"msghub"
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
	conns := [3]net.Conn{dial(port, t), dial(port, t), dial(port, t)}
	var user *main.User

	for i, conn := range conns {
		wantedId := i + 1
		sendPayload(conn, payload, t)
		unmarshal(&user, conn, t)

		if user == nil {
			t.Error("Expected user to not be nil")
		} else if user.Id != uint64(wantedId) {
			t.Errorf("Expected user ID to be %d, but got %d", wantedId, user.Id)
		}

		logout(conn, t)
	}
}

func TestListMessage(t *testing.T) {
	payload := `{ "type": "getAllUsers" }`
	conns := [3]net.Conn{dial(port, t), dial(port, t), dial(port, t)}
	var users []*main.User

	sendPayload(conns[0], payload, t)
	unmarshal(&users, conns[0], t)

	if users == nil {
		t.Error("Expected users array to not be nil")
	} else if wanted, got := len(conns)-1, len(users); wanted != got {
		t.Errorf("Expected different users array length. Wanted %d, got %d", wanted, got)
	}

	for _, user := range users {
		if user.Id != 2 && user.Id != 3 {
			t.Error("User has incorrect ID:", user.Id)
		}
	}

	for _, conn := range conns {
		logout(conn, t)
	}
}

func TestListMessageRemovesLoggedOutUsers(t *testing.T) {
	payload := `{ "type": "getAllUsers" }`
	masterConn := dial(port, t)
	conns := [2]net.Conn{dial(port, t), dial(port, t)}

	for i := 0; i < len(conns) - 1; i++ {
		var users []*main.User
		sendPayload(masterConn, payload, t)
		unmarshal(&users, masterConn, t)

		if users == nil {
			t.Error("Expected users array to not be nil")
		} else if wanted, got := len(conns) - i, len(users); wanted != got {
			t.Errorf("Expected different users array length. Wanted %d, got %d", wanted, got)
		}

		logout(conns[i], t)
	}

	logout(masterConn, t)
}

func TestRelayMessage(t *testing.T) {
	payload := `{ "type": "sendMessage", "userIds": [5, 8], "message": "Hello" }`
	masterConn := dial(port, t)
	connsCount := 1
	conns := make([]net.Conn, connsCount)

	for i := 0; i < connsCount; i++ {
		conns[i] = dial(port, t)
	}

	sendPayload(masterConn, payload, t)

	for i := 0; i < connsCount; i++ {
		var message *main.UserMessage
		var err *main.ErrorMessage

		conn := conns[i]
		isRecipient := i == 7 || i == 8
		decoder := json.NewDecoder(conn)

		if (isRecipient) {
			decoder.Decode(&message)
		} else {
			decoder.Decode(&err)
		}

		if isRecipient && message == nil {
			t.Error("Intended recipient didn't receive message")
		} else if wanted, got := "Hello", message.Message; isRecipient && wanted != got {
			t.Errorf("Wrong contents transmitted. Got %s, wanted %s", got, wanted)
		} else if !isRecipient && err == nil {
			t.Error("Incorrect recipient received message")
		}

		logout(conn, t)
	}

	logout(masterConn, t)
}

func TestRelayMessageInvalidUser(t *testing.T) {
	payload := `{ "type": "sendMessage", "userIds": [10], "message": "Hello" }`
	conn := dial(port, t)

	sendPayload(conn, payload, t)

	var message *main.ErrorMessage
	unmarshal(&message, conn, t)

	if message == nil {
		t.Error("Expected error response to not be nil")
	} else if wanted, got := "User(s) not found", message.Message; wanted != got {
		t.Errorf("Wrong contents transmitted. Got %s, wanted %s", got, wanted)
	}

	logout(conn, t)
}

func TestInvalidCommand(t *testing.T) {
	payload := `{ "type": "foobar" }`
	conn := dial(port, t)
	var err *main.ErrorMessage

	sendPayload(conn, payload, t)
	unmarshal(&err, conn, t)

	if err == nil {
		t.Error("Expected error response to not be nil")
	} else if got, wanted := err.Message, "Command not found"; got != wanted {
		t.Error("Incorrect error message", wanted, got)
	}

	logout(conn, t)
}

func dial(port string, t *testing.T) net.Conn {
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

func logout(conn net.Conn, t *testing.T) {
	payload := `{ "type": "logout" }`
	sendPayload(conn, payload, t)
	var info *main.InfoMessage
	unmarshal(&info, conn, t)
	conn.Close()
}
