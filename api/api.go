package api

import (
	"fmt"
	"mygoproject/repository"
	"net/http"
)

func (s server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := repository.User{
		Name:  "John",
		Email: "John@gmail",
		Age:   20,
	}

	err := s.db.CreateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully created user")

	w.Write([]byte("Successfully created user"))
}
