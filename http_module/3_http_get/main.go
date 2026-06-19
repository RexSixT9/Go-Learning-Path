package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func fetchTodo() (Todo, error) {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		return Todo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d", resp.StatusCode)
		return Todo{}, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body) // Use io.ReadAll to read the response body
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
		return Todo{}, err
	}

	var data Todo
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v", err)
		return Todo{}, err
	}
	return data, nil
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	todo, err := fetchTodo()
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, http.StatusOK, todo)
}

func main() {
	http.HandleFunc("/fetch", fetchHandler)
	http.ListenAndServe(":3000", nil)
}
