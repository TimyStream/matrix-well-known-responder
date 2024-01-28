package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type jResponse struct {
	URL string `json:"m.server"`
}

func ts_responseHandler(w http.ResponseWriter, r *http.Request) {
	var encoded jResponse
	encoded.URL = os.Getenv("MATRIX_SERVER_URL")
	j, _ := json.Marshal(encoded)
	w.Write(j)
}

func main() {
	http.HandleFunc("/", ts_responseHandler)
	log.Println("HTTP Handler Started")
	http.ListenAndServe(":8080", nil)
}
