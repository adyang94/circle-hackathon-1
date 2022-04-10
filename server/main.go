package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/adyang94/react-go-todo-app/router"
)

func main () {
	r := router.Router()
	fmt.Println("starting the server on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}