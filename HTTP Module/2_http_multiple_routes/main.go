package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// writeJSON is a helper function to write JSON responses with a given status code and data.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// TestRequest represents the expected structure of the JSON request body for the /test endpoint.
type TestRequest struct {
	Name string `json:"name"`
}

// testHandler handles POST requests to the /test endpoint. It validates the request method and JSON body, returning appropriate responses.
func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}
	defer r.Body.Close()

	var req TestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid JSON",
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Name is required",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":   true,
		"data": req.Name,
	})
}

// healthHandler handles GET requests to the /health endpoint, returning a JSON response indicating the service is healthy along with the current UTC date and time.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":       true,
		"message":  "Success",
		"dateTime": time.Now().UTC(),
	})
}

// helloHandler handles GET requests to the /hello endpoint, returning a JSON response with a greeting message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":   true,
		"data": "Hello, " + name + "!",
	})
}

// rootHandler handles GET requests to the root endpoint ("/"), returning a JSON response with a welcome message.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":   true,
		"data": "Welcome to the root page!",
	})
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
