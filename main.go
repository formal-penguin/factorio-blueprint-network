package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
