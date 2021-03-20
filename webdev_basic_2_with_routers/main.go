package main

import (
	"fmt"
	"net/http"
)

/*

-----------------------
http.NewServeMux()
----------------------
http.NewServeMux() it just a function that returns :
	an object of a struct which implemented handle interface

*/

type number1 int

// Implement handle interface
func (p number1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi num1"))
}

type number2 int

// Implement handle interface
func (p number2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi num2"))
}

const port = ":3885"

func main() {
	mux := http.NewServeMux()
	var n1 number1
	var n2 number2

	mux.Handle("/num1", n1)
	mux.Handle("/num2", n2)

	fmt.Println("server running on  http://localhost" + port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
