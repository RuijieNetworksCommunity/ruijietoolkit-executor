package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shelltool/shelltool/constant"
)

// getAPIURL 根据 AppType 硬编码环境地址
func getAPIURL() string {
	switch constant.AppType {
	case "test":
		return "https://api.colin.1tip.cc"
	case "dev":
		return "http://localhost:8000"
	default:
		return "https://api.emtips.net"
	}
}

// VerifyLicenseKey 向后端 API 验证 Key 的有效性
func VerifyLicenseKey(accessToken string, licenseKey string) (bool, string) {
	fullURL := getAPIURL() + "/ruijie/executor/key/verify"
	log.Printf("[Auth] Debugging Token: Length=%d, Start=%s, End=%s",
		len(accessToken), accessToken[:5], accessToken[len(accessToken)-5:])

	reqBody := VerifyRequest{Key: licenseKey}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("[Auth] JSON Marshal Error: %v", err)
		return false, "Internal JSON Error"
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("[Auth] NewRequest Failed: %v", err)
		return false, "System Error"
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: constant.APITimeout}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Auth] Request Failed: %v", err)
		return false, "Auth Service Unavailable"
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("[Auth] Warning: Failed to close response body: %v", err)
		}
	}()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Auth] Read Response Body Failed: %v", err)
		return false, "Failed to read auth response"
	}

	var apiResp APIResponse
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		logSnippet := string(bodyBytes)
		if len(logSnippet) > 200 {
			logSnippet = logSnippet[:200] + "..."
		}
		log.Printf("[Auth] Parse JSON Error: %v. Status: %d, Body Snippet: %s", err, resp.StatusCode, logSnippet)
		return false, "Invalid Response from Auth Server"
	}

	if resp.StatusCode == http.StatusOK && apiResp.Code == 2002450 {
		return true, apiResp.Msg
	}

	log.Printf("[Auth] Verification Denied. Key: %s, HTTP: %d, Code: %d, Msg: %s",
		licenseKey, resp.StatusCode, apiResp.Code, apiResp.Msg)

	return false, apiResp.Msg
}
