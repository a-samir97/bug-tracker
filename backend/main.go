package main

import (
	"BugTracker/helpers"
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value(1000))

	w.Write([]byte("Hello World"))
}

func main() {
	_, err := helpers.InitDb()

	if err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/", helloWorld)
	fmt.Println("Database connected successfully.")

	http.ListenAndServe(":8000", nil)
}
