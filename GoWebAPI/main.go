package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//var todosRequestCount uint64

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		cancel()
	}()

	go func() {
		if err := startHTTPServer(); err != nil {
			log.Fatalf("Server failed : %v", err)
		}
	}()

	<-ctx.Done()
}

func startHTTPServer() (err error) {
	// Wrap the entire function in a defer and recover block
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	http.HandleFunc("/todos", todosHandler)
	//http.HandleFunc("/todosCounter", todosCountHandler)
	log.Println("Starting server on port 8080...")
	return http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(GetToDos()); err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
	}

	//atomic.AddUint64(&todosRequestCount, 1)
	//currentCount := atomic.LoadUint64(&todosRequestCount)
	//log.Printf("Todos endpoint called %d times", currentCount)
}

// func todosCountHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	currentCount := atomic.LoadUint64(&todosRequestCount)
// 	if err := json.NewEncoder(w).Encode(currentCount); err != nil {
// 		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
// 	}

// 	log.Printf("Todos endpoint called %d times", currentCount)
// }
