package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func doRequest(client *http.Client, method, url string, body []byte, token string) {
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		fmt.Printf("ERROR building request %s %s: %v\n", method, url, err)
		return
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR request %s %s: %v\n", method, url, err)
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n==> %s %s\nStatus: %d\nBody:\n%s\n", method, url, resp.StatusCode, string(b))
}

func login(client *http.Client, base, email, pass string) (string, map[string]interface{}) {
	payload := map[string]string{"email": email, "password": pass}
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", base+"/api/auth/login", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("login error: %v\n", err)
		return "", nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var out map[string]interface{}
	_ = json.Unmarshal(body, &out)
	token, _ := out["token"].(string)
	return token, out
}

func main() {
	base := os.Getenv("API_BASE")
	if base == "" {
		base = "http://localhost:8080"
	}
	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Println("Base:", base)

	// health
	doRequest(client, "GET", base+"/api/", nil, "")

	// list products
	doRequest(client, "GET", base+"/api/products", nil, "")

	// sample product
	sample := "ab6cc690-de98-4642-a027-9e04728f719d"
	doRequest(client, "GET", base+"/api/products/"+sample, nil, "")

	// login user
	userToken, userResp := login(client, base, "demo@local.com", "password")
	fmt.Printf("\nUSER LOGIN RESPONSE: %+v\n", userResp)

	// login admin
	adminToken, adminResp := login(client, base, "admin@local.com", "adminpass")
	fmt.Printf("\nADMIN LOGIN RESPONSE: %+v\n", adminResp)

	// protected endpoints
	doRequest(client, "GET", base+"/api/user/orders", nil, userToken)
	doRequest(client, "GET", base+"/api/admin/products", nil, adminToken)

	// try create product (admin)
	prod := map[string]interface{}{"name": "Smoke Product", "price": 9.5, "stock": 3}
	pb, _ := json.Marshal(prod)
	doRequest(client, "POST", base+"/api/admin/products", pb, adminToken)
}
