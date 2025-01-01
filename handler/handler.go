package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}

func NewUserHandler(name, email string, age int) *User {
	return &User{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func UserNameHandler(user User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(user.Name))
	}
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.Age = u.Age + 1
	w.Write([]byte(u.Name + ": " + fmt.Sprint(u.Age)))
}
