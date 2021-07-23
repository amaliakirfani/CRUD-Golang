package main

import (
	"AttendanceApi/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	route := routes.Routes()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var currentTime = time.Now()
	var curTime = currentTime.Format("2006-01-02 15:04:05")

	fmt.Println(curTime)

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", route))
}
