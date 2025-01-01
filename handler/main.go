package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	user := NewUserHandler("kunkka", "kunkka@gmail.com", 20)
	mux.Handle("/age", user)
	mux.HandleFunc("/name", UserNameHandler(*user))
	fmt.Println(http.ListenAndServe(":8080", mux))
}
