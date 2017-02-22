package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const DEFAULT_PORT int = 0

var (
	cGRAPH = flag.String("g", "", "path to graph")
	cPORT  = flag.Int("p", DEFAULT_PORT, "Port to listen on")
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(",")))

	s := &http.Server{
		Handler:        &loggingMiddleware{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *cPORT))
	if err != nil {
		log.Fatalf("can't bind server: %s", err)
	}
	// output the address we bound
	log.Printf("bound address: %s", ln.Addr().String())
	log.Fatal(s.Serve(ln))
}

type loggingMiddleware struct {
}

func (*loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	http.DefaultServeMux.ServeHTTP(w, r)
	log.Printf("%s [%s] %s", r.Method, r.URL.Path, time.Since(startTime))
}
