package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"shelltool/shelltool/constant"
	"time"

	"github.com/beevik/ntp"
)

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

// syncTime 自动同步时间
func syncTime() {
	for _, ntpServer := range constant.NTPServers {
		ntpTime, err := ntp.Time(ntpServer)
		if err == nil {
			log.Printf("Time synchronized with %s: %s", ntpServer, ntpTime.Format(time.RFC1123))
			return
		}
	}
	log.Printf("Failed to sync time with any NTP server.")
}

func main() {
	syncTime()
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
