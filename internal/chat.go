package internal

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	Clients   map[*websocket.Conn]string
	Broadcast chan string
	Mu        sync.Mutex
}

var chatRoom = &ChatRoom{
	Clients:   make(map[*websocket.Conn]string),
	Broadcast: make(chan string),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
