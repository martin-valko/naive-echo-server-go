/*
  Copyright (c) Martin Valkó. All rights reserved.
  Licensed under the MIT license; check the LICENSE file in the project root for details.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type EchoResponse struct {
	Path string `json:"path"`
}

const port string = "8080"

func echoHandler(w http.ResponseWriter, r *http.Request) {
	response := &EchoResponse{
		Path: r.URL.Path,
	}

	jsonBytes, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Processing request '%s'", response.Path)
	fmt.Fprint(w, string(jsonBytes))
}

func main() {
	log.Printf("Server is starting on port '%s'.", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", echoHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)

	if err != nil {
		log.Fatal(err)
	}
}
