package constant

import "time"

const (
	Version = "0.0.1-nekoapi-ruijietoolkit-executor"

	// AppType 应用程序类型
	AppType = "test"

	// ServerAddr 执行器监听端口
	ServerAddr = "0.0.0.0:54134"

	// APITimeout 请求后端超时时间
	APITimeout = 5 * time.Second
)

var (
	// DefaultShell 默认 Shell
	DefaultShell = "/bin/sh"
)
