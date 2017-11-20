package main

import (
	"bufio"
	"flag"
	"net/http"
	"os"
	"time"

	log "github.com/golang/glog"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 3 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	log.Info("Start listen on :8080")
	http.HandleFunc("/apis/v1/logstream", serveWs)
	http.ListenAndServe(":8080", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs upgrades the HTTP server connection to the WebSocket protocol.
// Writes the message to WebSocket connection.
func serveWs(rw http.ResponseWriter, req *http.Request) {
	log.Infof("Upgrade to websocket")
	ws, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}

	writer(ws)
}

// writer reads the message from standard input, and writes it into the WebSocket connection.
func writer(ws *websocket.Conn) {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		ws.Close()
	}()

	// Read the message from standard input.
	lines := make(chan []byte, 10)
	go func() {
		for {
			input := bufio.NewReader(os.Stdin)
			line, err := input.ReadSlice('\n')
			if err != nil {
				log.Fatalf("input error: %s", err)
			}

			lines <- line
		}
	}()

	for {
		select {
		case line, ok := <-lines:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := ws.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
				log.Error(err.Error())
				return
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Error(err.Error())
				return
			}
		}
	}
}
