package main

import (
	"log"
	"net/http"
	"testing01/handlers"
)

func main() {
	users := map[string]handlers.User{
		"user1": {
			Name:     "Test",
			LastName: "Test",
		},
		"user2": {
			Name:     "Test 2",
			LastName: "Test 2",
		},
	}
	http.HandleFunc("/user", handlers.UserViewHandler(users))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
