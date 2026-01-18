package constant

import "time"

var (
	AppName = "nekoapi-ruijietoolkit-executor"

	Version = "0.0.0-self_build"

	BuildTime = "self_build"

	// AppType 应用程序类型
	AppType = "dev"

	// ServerAddr 执行器监听端口
	ServerAddr = "0.0.0.0:54134"

	// APITimeout 请求后端超时时间
	APITimeout = 5 * time.Second

	// DefaultShell 默认 Shell
	DefaultShell = "/bin/sh"
)
