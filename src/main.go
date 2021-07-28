package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const ServerAddr = "127.0.0.1"
const ServerPort = "8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create_index", CreateIndexHandler).Methods("POST")
	r.HandleFunc("/insert", DataInsertHandler).Methods("POST")
	r.HandleFunc("/delete", DataDeleteHandler).Methods("POST")
	r.HandleFunc("/search", DataSearchHandler).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", ServerAddr, ServerPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Printf("Server running: %s:%s\n", ServerAddr, ServerPort)
	log.Fatal(srv.ListenAndServe())
}
