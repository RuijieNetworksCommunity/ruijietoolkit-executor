package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"shelltool/shelltool/constant"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	var token = r.URL.Query().Get("token")
	var licences = r.URL.Query().Get("licences")
	if len(token) <= 0 || len(licences) <= 0 {
		// w.Write([]byte("need token and licences"))
		return
	}

	is_valid := get_licences_is_valid(token, licences)
	if !is_valid {
		// w.Write([]byte("licences invalid"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		// w.Write([]byte("something wrong"))
		return
	}
	defer conn.Close()

	var cmd *exec.Cmd

	switch constant.Type {
	case "dev":
		cmd = exec.Command("/bin/login")
	case "localdev":
		cmd = exec.Command("/bin/bash")
	default:
		cmd = exec.Command("/bin/login")
	}

	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Println("pty:", err)
		// w.Write([]byte("something wrong"))
		return
	}
	defer ptmx.Close()

	// PTY → WebSocket
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				return
			}
			conn.WriteMessage(websocket.BinaryMessage, buf[:n])
		}
	}()

	// WebSocket → PTY
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// w.Write([]byte("something wrong"))
			return
		}
		ptmx.Write(msg)
	}
}

func main() {
	fmt.Printf("Version: %s Buildtime: %s\n", constant.Version, constant.BuildTime)

	http.HandleFunc("/ws", wsHandler)

	// 之后会移除
	if constant.Type != "release" {
		http.Handle("/", http.FileServer(http.Dir("./static")))
	}

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
