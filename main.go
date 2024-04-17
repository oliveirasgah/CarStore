package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    server := &http.Server {
        Addr: "localhost:8000",
        ReadTimeout: time.Second * 10,
        WriteTimeout: time.Second * 10,
    }

    fmt.Printf("Listening on %v\n", server.Addr)
    server.ListenAndServe()
}
