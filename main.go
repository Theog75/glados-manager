package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", controllers.Ping())

	fmt.Println("Glados Manager starting to listen on port " + config.Port)
	http.ListenAndServe(":"+config.Port, nil)
}
