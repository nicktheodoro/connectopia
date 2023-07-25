package main

import (
	"connectopia-api/src/config"
	"connectopia-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println(config.Port)
	fmt.Println("Api running")

	r := router.Generate()

	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
