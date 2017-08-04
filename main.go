package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	if err := enc.Encode(data); err != nil {
		log.Println("encode response: %s", err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, struct {
			RemoteAddr string
			Method     string
			RequestURI string
			Proto      string
			Host       string
			Headers    http.Header
		}{
			RemoteAddr: r.RemoteAddr,
			Method:     r.Method,
			RequestURI: r.RequestURI,
			Proto:      r.Proto,
			Host:       r.Host,
			Headers:    r.Header,
		})
	})
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
