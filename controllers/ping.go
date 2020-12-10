package controllers

import (
	"encoding/json"
	"net/http"
)

//Ping retreive health checks
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Activated Ping test")
		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
