package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type messageResponse struct {
	Code int    `json:"code" xml:"code"`
	Msg  string `json:"msg" xml:"msg"`
	Data any    `json:"data,omitempty" xml:"data,omitempty"`
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// DefaultServeWs ServeWs handles websocket requests from the peer.
func DefaultServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) error {
	client, err := InitClient(hub, w, r)
	if err != nil {
		return err
	}

	Run(client)
	return nil
}

func InitClient(hub *Hub, w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		id:   time.Now().UnixNano(),
		hub:  hub,
		conn: conn,
		send: make(chan []byte, bufSize),
		readHandle: func(message []byte) []byte {
			return message
		},
		writeHandle: func(message []byte) []byte {
			return message
		},
	}

	return client, nil
}

func Run(client *Client) {
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
