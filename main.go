package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	interfaces "test-stone/Interfaces"
)

func main() {
	os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	Run(port)
}

func Run(port string) error {
	log.Printf("Server running at http://localhost:%s/", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), interfaces.Routes())
}
