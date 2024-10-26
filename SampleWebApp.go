package main

import (
	"fmt"
	"net/http"
	_ "net/http"
)

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-user", displayUser)

	http.ListenAndServe(":7777", nil)

}

func displayUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello Golang User ! How are you ?"
	var name string = "Shivananda"
	var age string = "35"
	var gender string = "Gender"
	fmt.Fprintf(writer, greeting, name, age, gender)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello Golang User ! How are you ?"
	fmt.Fprintf(writer, greeting)
}
