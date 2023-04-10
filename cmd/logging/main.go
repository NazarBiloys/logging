package main

import (
	"fmt"
	"github.com/NazarBiloys/mysql-banchmarks/internal/service"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := service.MakeUser(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
}

func handlerRead(w http.ResponseWriter, r *http.Request) {
	if err := service.GetUser(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", handlerRead)
	log.Fatal(http.ListenAndServe(":90", nil))
}
