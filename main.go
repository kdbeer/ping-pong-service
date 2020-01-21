package main

import (
	"fmt"
	"net/http"
	"time"
)

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[pingPongHandler] request at", time.Now())
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", pingPongHandler)

	fmt.Println("ping listening on 0.0.0.0, port 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error starting ping server: ", err)
	}
}
