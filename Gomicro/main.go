package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/person", personHandler)

	http.ListenAndServe("localhost:8082", nil)
}

type Person struct {
	Name string
	Age  int
}

type Response struct {
	Success bool
	Data    any
	Error   string
}

func personHandler(asdfasdf http.ResponseWriter, qwerqwer *http.Request) {

	switch qwerqwer.Method {
	case http.MethodGet:
		persons := make([]Person, 0)

		persons = append(persons, Person{
			Name: "Budi",
			Age:  10,
		})

		persons = append(persons, Person{
			Name: "Ani",
			Age:  20,
		})

		var r Response = Response{
			Success: true,
			Data:    persons,
		}

		foo(r, asdfasdf)
	default:
		var r Response = Response{
			Success: false,
			Error:   "not found",
		}

		foo(r, asdfasdf)
	}

}

func foo(r Response, w http.ResponseWriter) {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error marshalling", err)
		http.Error(w, fmt.Sprintf("error json encoding %s", err), http.StatusOK)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ini dari fmt.Println") // io server
	w.Write([]byte("hello world"))
}
