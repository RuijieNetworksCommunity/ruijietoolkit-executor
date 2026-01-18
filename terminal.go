package main

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"shelltool/shelltool/constant"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

// StartInteractiveShell 启动交互式终端会话
func StartInteractiveShell(conn *websocket.Conn, initMsg SessionInit) {
	log.Printf("[Shell] Starting PTY Session (Rows: %d, Cols: %d)", initMsg.Rows, initMsg.Cols)

	c := exec.Command(constant.DefaultShell)

	ptmx, err := pty.StartWithSize(c, &pty.Winsize{
		Rows: initMsg.Rows,
		Cols: initMsg.Cols,
	})
	if err != nil {
		log.Printf("[Shell] Failed to start PTY: %v", err)

		errMsg := map[string]string{"error": "Failed to start PTY: " + err.Error()}
		if writeErr := conn.WriteJSON(errMsg); writeErr != nil {
			log.Printf("[Shell] Failed to send error to client: %v", writeErr)
		}
		return
	}

	defer func() {
		if err := ptmx.Close(); err != nil {
			// PTY 关闭错误通常是由于进程已经结束导致的，记录 Debug 级别日志即可
			// log.Printf("[Shell] PTY Close Error: %v", err)
		}

		if c.Process != nil {
			if err := c.Process.Kill(); err != nil {
				// 进程可能已经自己退出了，忽略错误
			}
		}
		log.Println("[Shell] Session ended")
	}()

	go func() {
		buf := make([]byte, 4096) // 4KB 缓冲
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				if err != io.EOF {
					// Linux 下 EIO 错误通常意味着所有 PTY 句柄都关闭了（即 Shell 退出了）
					// 这是一个正常的退出信号，非严重错误
					// log.Printf("[Shell] PTY Read Error: %v", err)
				}
				return
			}
			if err := conn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
				// log.Printf("[Shell] WS Write Error: %v", err)
				return
			}
		}
	}()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			// 前端断开连接
			break
		}

		// 情况 A: 二进制消息 -> 直接当作 STDIN
		if mt == websocket.BinaryMessage {
			// [FIX] 处理 PTY Write 错误
			// 如果写入 PTY 失败，说明 Shell 进程可能已经挂了，应退出循环
			if _, err := ptmx.Write(message); err != nil {
				log.Printf("[Shell] PTY Write Error: %v", err)
				break
			}
			continue
		}

		var input ShellInput
		if json.Unmarshal(message, &input) == nil && input.Type != "" {
			switch input.Type {
			case "resize":
				// 调整窗口失败通常不致命，记录日志即可，不要 break
				if err := pty.Setsize(ptmx, &pty.Winsize{Rows: input.Rows, Cols: input.Cols}); err != nil {
					log.Printf("[Shell] Resize Error: %v", err)
				}
			case "stdin":
				if _, err := ptmx.Write([]byte(input.Data)); err != nil {
					log.Printf("[Shell] PTY Write (Stdin) Error: %v", err)
					break // 退出循环
				}
			}
			continue
		}

		if _, err := ptmx.Write(message); err != nil {
			log.Printf("[Shell] PTY Write (Text) Error: %v", err)
			break
		}
	}
}
