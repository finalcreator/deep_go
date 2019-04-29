package simpleserver

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendjson", func(w http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "mike",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	})
}
