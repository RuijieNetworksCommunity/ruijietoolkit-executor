package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"shelltool/shelltool/constant"
)

// =======================
// 全局配置
// =======================

// findAvailableShell 自动探测系统可用的 Shell
func findAvailableShell() string {
	shells := []string{"/bin/bash", "/bin/ash", "/usr/bin/bash", "/bin/zsh", "/bin/sh"}

	for _, s := range shells {
		if info, err := os.Stat(s); err == nil && info.Mode()&0111 != 0 {
			return s
		}
		if path, err := exec.LookPath(s); err == nil {
			return path
		}
	}
	return "/bin/sh"
}

func main() {
	log.Printf("Version: %s-%s", constant.Version, constant.AppName)
	log.Printf("BuildTime: %s", constant.BuildTime)
	constant.DefaultShell = findAvailableShell()
	log.Printf("Detected Shell: %s", constant.DefaultShell)

	http.HandleFunc("/ws", HandleWebSocket)

	log.Printf("Executor Started | Mode: %s | API: %s", constant.AppType, getAPIURL())

	if err := http.ListenAndServe(constant.ServerAddr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
