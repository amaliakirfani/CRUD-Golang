package main

import (
	"AttendanceApi/routes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	route := routes.Routes()

	var currentTime = time.Now()
	var curTime = currentTime.Format("15:04:05")

	if curTime < "17:00:00" {
		fmt.Println("blm jam 17")
	}

	// var layoutFormat, value string
	// var date time.Time

	// layoutFormat = "15:04:05"
	// value = "08:00:00"
	// date, _ = time.Parse(layoutFormat, value)
	// fmt.Println(value, "\t->", date.String())

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", route))
}
