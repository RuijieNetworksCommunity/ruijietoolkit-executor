package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // 允许跨域
}

// HandleWebSocket 处理主入口
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("access_token")
	if token == "" {
		http.Error(w, "Unauthorized: Missing access_token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade Error: %v", err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			// 连接关闭时的错误通常记录为 Debug/Info 级别，因为这是正常的网络行为
			// log.Printf("WS Close: %v", err)
		}
	}()

	if !performHandshake(conn, token) {
		return // 认证失败，关闭连接
	}

	initMsg, err := readInitMessage(conn)
	if err != nil {
		log.Printf("[Protocol] Failed to read init message: %v", err)
		return
	}

	switch initMsg.Type {
	case "shell":
		StartInteractiveShell(conn, *initMsg)
	case "exec":
		StartCommandExecution(conn, initMsg.Command)
	default:
		errMsg := map[string]string{"error": "Unknown session type"}
		if err := conn.WriteJSON(errMsg); err != nil {
			log.Printf("[Protocol] Failed to send error message: %v", err)
		}
	}
}

// 内部辅助：执行握手验证
func performHandshake(conn *websocket.Conn, token string) bool {
	if err := conn.WriteJSON(map[string]string{"type": "AUTH_REQUIRED", "msg": "Please send license key"}); err != nil {
		log.Printf("[Handshake] Write AUTH_REQUIRED failed: %v", err)
		return false
	}

	if err := conn.SetReadDeadline(time.Now().Add(30 * time.Second)); err != nil {
		log.Printf("[Handshake] SetReadDeadline failed: %v", err)
		return false
	}

	_, keyBytes, err := conn.ReadMessage()
	if err != nil {
		log.Printf("[Handshake] Read key failed: %v", err)
		return false
	}

	isValid, msg := VerifyLicenseKey(token, string(keyBytes))

	if !isValid {
		resp := map[string]interface{}{"type": "AUTH_FAILED", "success": false, "msg": msg}
		if err := conn.WriteJSON(resp); err != nil {
			log.Printf("[Handshake] Write AUTH_FAILED failed: %v", err)
		}
		return false
	}

	successResp := map[string]interface{}{"type": "AUTH_SUCCESS", "success": true, "msg": "Authorized"}
	if err := conn.WriteJSON(successResp); err != nil {
		log.Printf("[Handshake] Write AUTH_SUCCESS failed: %v", err)
		return false
	}

	if err := conn.SetReadDeadline(time.Time{}); err != nil {
		log.Printf("[Handshake] Clear ReadDeadline failed: %v", err)
		return false
	}

	return true
}

// 内部辅助：读取初始化指令
func readInitMessage(conn *websocket.Conn) (*SessionInit, error) {
	_, msgBytes, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	var initMsg SessionInit
	if err := json.Unmarshal(msgBytes, &initMsg); err != nil {
		return nil, err
	}
	return &initMsg, nil
}
