package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {

		resp := map[string]string{
			"name": "Naman",
		}

		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":9000", nil)
}