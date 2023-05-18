package main

import (
	"fmt"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

// ...
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		panic(err)
	}

	// TODO: remove
	var user *User
	fmt.Println(user.Name)
}
