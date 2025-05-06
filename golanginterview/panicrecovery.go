package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime/debug"
	"time"
)

func mainy() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/random", RandomHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: RecoveryMiddleware(mux),
	}
	log.Fatal(server.ListenAndServe())
}
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(10)
	if n%2 == 0 {
		panic("something bad")
	}
	fmt.Fprintf(w, "Hello")
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic : %v\n %s", rec, debug.Stack())
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

/*
func main() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/random", RandomHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: RecoverMiddleware(mux), // why are we sending mux here ?
	}
	log.Println("🚀 API running at http://localhost:8080/random")
	log.Fatal(server.ListenAndServe())
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("PANIC: %v\n%s", rec, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(10)
	if n%3 == 0 {
		panic("something bad")
	}
	fmt.Fprintf(w, "success %d", n)
}

*/
