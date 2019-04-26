package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "user login")
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", Login)
	err := http.ListenAndServe("0.0.0.0:8889", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}

}
