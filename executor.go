package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"shelltool/shelltool/constant"

	"github.com/gorilla/websocket"
)

// StartCommandExecution 执行非交互式命令
func StartCommandExecution(conn *websocket.Conn, cmdStr string) {
	log.Printf("[Exec] Command: %s", cmdStr)

	// 使用 sh -c 允许复杂的命令链 (例如: "cd /var && ls -la")
	cmd := exec.Command(constant.DefaultShell, "-c", cmdStr)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	result := ExecResult{
		Type:    "EXEC_RESULT",
		Command: cmdStr,
		Stdout:  stdoutBuf.String(),
		Stderr:  stderrBuf.String(),
		Code:    0,
	}

	if err != nil {
		result.Code = -1

		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			result.Code = exitErr.ExitCode()
		} else {
			// 如果不是 ExitError (例如路径找不到 exec: not found)，打印日志
			log.Printf("[Exec] Generic Error: %v", err)
		}
	}

	if err := conn.WriteJSON(result); err != nil {
		log.Printf("[Exec] Error sending result to client: %v", err)
	}
	
	if err := conn.Close(); err != nil {
		// 记录 Warning，不影响流程
		log.Printf("[Exec] Warning: Failed to close WebSocket connection: %v", err)
	}

	log.Printf("[Exec] Finished. Exit Code: %d", result.Code)
}
