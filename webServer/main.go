package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//This function accepts two parameters:
	//The first one is a string representing the path to match and the second a function that will be executed to handle the request.
	// http.ResponseWriter handles any response
	//http.Request is used to get any data or value from the request

	http.HandleFunc("/users", usersHandleFunc)
	//os.Exit(2)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func usersHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("We got a request on /users")
	fmt.Fprintf(w, "Hi, thanks for calling my /users API with HTTP method '%v'", r.Method)
}
