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

func doRequest(client *http.Client, method, url string, body []byte, token string) (int, []byte) {
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
		return 0, nil
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR request %s %s: %v\n", method, url, err)
		return 0, nil
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n==> %s %s\nStatus: %d\nBody:\n%s\n", method, url, resp.StatusCode, string(b))
	return resp.StatusCode, b
}

func login(client *http.Client, base, email, pass string) (string, map[string]interface{}) {
	payload := map[string]string{"email": email, "password": pass}
	b, _ := json.Marshal(payload)
	_, body := doRequest(client, "POST", base+"/api/auth/login", b, "")
	var out map[string]interface{}
	_ = json.Unmarshal(body, &out)
	tok, _ := out["token"].(string)
	return tok, out
}

func main() {
	base := os.Getenv("API_BASE")
	if base == "" {
		base = "http://localhost:8080"
	}
	client := &http.Client{Timeout: 15 * time.Second}

	fmt.Println("Base:", base)

	// health
	doRequest(client, "GET", base+"/api/", nil, "")

	// register a temp user
	ts := time.Now().Unix()
	reg := map[string]string{"name": "SmokeUser", "email": fmt.Sprintf("smoke-%d@local.com", ts), "password": "password"}
	rb, _ := json.Marshal(reg)
	doRequest(client, "POST", base+"/api/auth/register", rb, "")

	// login user
	userEmail := reg["email"]
	userToken, _ := login(client, base, userEmail, "password")

	// auth/me
	doRequest(client, "GET", base+"/api/auth/me", nil, userToken)

	// list products
	_, productsBody := doRequest(client, "GET", base+"/api/products", nil, "")
	var productsResp map[string]interface{}
	_ = json.Unmarshal(productsBody, &productsResp)
	var sampleID string
	if items, ok := productsResp["items"].([]interface{}); ok && len(items) > 0 {
		if itm, ok := items[0].(map[string]interface{}); ok {
			if id, ok := itm["id"].(string); ok {
				sampleID = id
			}
		}
	}

	// admin login
	adminToken, _ := login(client, base, "admin@local.com", "adminpass")

	// admin create product
	newProd := map[string]interface{}{"name": "FS-SmokeProd", "price": 5.5, "stock": 10}
	npb, _ := json.Marshal(newProd)
	_, newProdBody := doRequest(client, "POST", base+"/api/admin/products", npb, adminToken)
	var newProdResp map[string]interface{}
	_ = json.Unmarshal(newProdBody, &newProdResp)
	var newProdID string
	if id, ok := newProdResp["id"].(string); ok {
		newProdID = id
	}

	// admin update product
	if newProdID != "" {
		upd := map[string]interface{}{"name": "FS-SmokeProd-Updated", "price": 6.0, "stock": 8}
		ub, _ := json.Marshal(upd)
		doRequest(client, "PUT", base+"/api/admin/products/"+newProdID, ub, adminToken)
	}

	// admin list products
	doRequest(client, "GET", base+"/api/admin/products", nil, adminToken)

	// admin list users
	_, usersBody := doRequest(client, "GET", base+"/api/admin/users", nil, adminToken)
	var users []interface{}
	_ = json.Unmarshal(usersBody, &users)
	var targetUserID string
	for _, u := range users {
		if um, ok := u.(map[string]interface{}); ok {
			if em, ok := um["email"].(string); ok && em == userEmail {
				if id, ok := um["id"].(string); ok {
					targetUserID = id
				}
			}
		}
	}

	// admin block user (send JSON body { blocked: true })
	if targetUserID != "" {
		blk, _ := json.Marshal(map[string]bool{"blocked": true})
		doRequest(client, "PUT", base+"/api/admin/users/"+targetUserID+"/block", blk, adminToken)
	}

	// create order as user (use sampleID or newProdID)
	orderProd := sampleID
	if orderProd == "" {
		orderProd = newProdID
	}
	if orderProd != "" {
		ord := map[string]interface{}{"items": []map[string]interface{}{{"id": orderProd, "qty": 1}}}
		ob, _ := json.Marshal(ord)
		_, orderBody := doRequest(client, "POST", base+"/api/user/orders", ob, userToken)
		var orderResp map[string]interface{}
		_ = json.Unmarshal(orderBody, &orderResp)
		var orderID string
		if id, ok := orderResp["id"].(string); ok {
			orderID = id
		}

		// get user orders
		doRequest(client, "GET", base+"/api/user/orders", nil, userToken)

		if orderID != "" {
			// get order by id
			doRequest(client, "GET", base+"/api/user/orders/"+orderID, nil, userToken)

			// user update order status
			doRequest(client, "PUT", base+"/api/user/orders/"+orderID+"/status", []byte("{\"status\":\"cancelled\"}"), userToken)

			// admin update order status (use allowed status: completed)
			doRequest(client, "PUT", base+"/api/admin/orders/"+orderID+"/status", []byte("{\"status\":\"completed\"}"), adminToken)
		}
	}

	// cleanup: admin delete created product
	if newProdID != "" {
		doRequest(client, "DELETE", base+"/api/admin/products/"+newProdID, nil, adminToken)
	}

	fmt.Println("Full smoke finished")
}
