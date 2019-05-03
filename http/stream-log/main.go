package main

import (
	"bufio"
	"flag"
	"net/http"
	"net/url"
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
	http.HandleFunc("/apis/v1/proxyws", proxyWs)
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
	log.Infof("rawquery: %+v", req.URL)
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

// proxyWs proxy the WebSocket request to another server.
// It will create two WebSocket connections and transfer data between them.
func proxyWs(rw http.ResponseWriter, req *http.Request) {
	clientConn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer clientConn.Close()

	// Proxy the request to another WebSocket handler of this server.
	u := url.URL{
		Host:   "127.0.0.1:8080",
		Path:   "/apis/v1/logstream",
		Scheme: "ws",
	}

	// Need to filter the header for upgrade, as these headers will be added when dial the server.
	// Otherwise, there will be duplicate header error.
	header := filterHeader(req.Header)

	serverConn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer serverConn.Close()

	go transfer(clientConn, serverConn)
	transfer(serverConn, clientConn)
}

// transfer transfers the data between two WebSocket connection.
func transfer(from *websocket.Conn, to *websocket.Conn) {
	for {
		typ, message, err := from.ReadMessage()
		if err != nil {
			log.Errorf("[websocket]: transfer from %v to %v, read error: %v", from.RemoteAddr(), to.RemoteAddr(), err)
			if closeErr := to.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, err.Error())); closeErr != nil {
				log.Errorf("[websocket] can't send close message to server: %v", closeErr)
			}
			return
		} else if err := to.WriteMessage(typ, message); err != nil {
			log.Errorf("[websocket]: transfer from %v to %v, write error: %v", from.RemoteAddr(), to.RemoteAddr(), err)
			if closeErr := from.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, err.Error())); closeErr != nil {
				log.Errorf("[websocket] can't send close message to client: %v", closeErr)
			}
			return
		}
	}
}

// filterHeader filters the headers for upgrading the HTTP server connection to the WebSocket protocol.
func filterHeader(header http.Header) http.Header {
	newHeader := http.Header{}
	for k, vs := range header {
		switch {
		case k == "Upgrade" ||
			k == "Connection" ||
			k == "Sec-Websocket-Key" ||
			k == "Sec-Websocket-Version" ||
			k == "Sec-Websocket-Extensions":
		default:
			newHeader[k] = vs
		}
	}
	return newHeader
}
