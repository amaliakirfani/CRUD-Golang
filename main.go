package main

import (
	"ToDoList/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	route := routes.Routes()

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", route))
}
