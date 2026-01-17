package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"shelltool/shelltool/constant"

	"net/http"
)

func get_api_url() string {
	switch constant.Type {
	case "dev":
		return "https://api.colin.1tip.cc"
	case "localtest":
		return "http://localhost:8000"
	default:
		return "https://api.emtips.net"
	}
}

func get_licences_is_valid(token string, licences string) bool {

	req_payload := map[string]string{
		"license_key": licences,
	}

	req_payload_bytes, _ := json.Marshal(req_payload)

	fmt.Println(string(req_payload_bytes))
	req, err := http.NewRequest("POST", get_api_url()+"/license/verify", bytes.NewBuffer(req_payload_bytes))
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	if resp.StatusCode != 200 {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	var jsonObj map[string]any
	err = json.Unmarshal(body, &jsonObj)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	if jsonObj["data"] == nil {
		return false
	}
	data_data := jsonObj["data"].(map[string]any)
	is_valid := data_data["valid"]

	if _, ok := is_valid.(bool); !ok {
		return false
	}

	return is_valid.(bool)
}
