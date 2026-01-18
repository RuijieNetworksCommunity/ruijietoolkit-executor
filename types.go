package main

// =======================
// 后端 API 相关结构
// =======================

// APIResponse 代表后端 API 返回的标准 JSON 响应格式
type APIResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// VerifyRequest 代表发送给后端 API 用于验证 License Key 的请求体
type VerifyRequest struct {
	Key string `json:"key"`
}

// =======================
// WebSocket 协议相关结构
// =======================

// SessionInit 代表客户端发送的初始化会话指令 (握手后第一条消息)
type SessionInit struct {
	Type    string `json:"type"`    // "shell" (交互式) 或 "exec" (单次执行)
	Command string `json:"command"` // 仅 exec 模式：要执行的命令
	Rows    uint16 `json:"rows"`    // 仅 shell 模式：窗口行数
	Cols    uint16 `json:"cols"`    // 仅 shell 模式：窗口列数
}

// ShellInput 代表 Shell 模式下客户端发送的操作指令 (输入或调整窗口)
type ShellInput struct {
	Type string `json:"type"` // "stdin" (输入) 或 "resize" (调整窗口)
	Data string `json:"data"` // 输入的内容 (Type为stdin时有效)
	Rows uint16 `json:"rows"` // resize 用的行 (Type为resize时有效)
	Cols uint16 `json:"cols"` // resize 用的列 (Type为resize时有效)
}

// ExecResult 代表 Exec 模式下命令执行完毕后返回的完整结果
type ExecResult struct {
	Type    string `json:"type"` // 固定为 "EXEC_RESULT"
	Command string `json:"command"`
	Stdout  string `json:"stdout"`
	Stderr  string `json:"stderr"`
	Code    int    `json:"code"` // 退出码 (0为成功)
}
