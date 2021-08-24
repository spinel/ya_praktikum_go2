package main

import (
	"SpinelChi/middleware"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// some long request here
	w.Write([]byte("request done"))
	time.Sleep(1 * time.Second)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.TimerTrace)
	r.Get("/", handler)
	log.Fatal(http.ListenAndServe(":8083", r))
}
