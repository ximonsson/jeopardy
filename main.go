package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("jeopardy")
	go Run() // start game
	log.Fatal(http.ListenAndServe(":8000", nil))
}
