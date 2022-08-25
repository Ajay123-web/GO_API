package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Ajay123-web/MONGODB/routers"
)

func main() {
	fmt.Println("Server Working!!!")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000" , r))
}