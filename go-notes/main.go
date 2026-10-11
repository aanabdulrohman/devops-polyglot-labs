package main

import(
	"encoding/json"
	"net/http"
)

type Note struct {
	ID			string `json:"id"`
	Content		string `json:"content"`
}

func main() {
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Note{ID: "1", Content: "This is a sample note."})
	})

	http.ListenAndServe(":8080", nil)
}