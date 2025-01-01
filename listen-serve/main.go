package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(http.ListenAndServe(":8080", &test{}))
}

type test struct{}

func (t *test) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("you are noob"))
}
