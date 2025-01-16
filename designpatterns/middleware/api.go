package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", auth(logger(handler)))
	http.ListenAndServe(":8080", nil)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	//http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hello from path %s", r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, msg)

}
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("i did auth")
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request is here")
		next(w, r)
	}
}
