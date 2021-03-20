package main

import (
	"fmt"
	"net/http"
)

/*

--------------------
ListenAndServe()
--------------------
it is a function.
It accepts two parameters
	1. 	Port with a colon like (:8023)
	2. 	An object of a user-defined type => that implements Handler interface.
		or
		Sometimes we pass nil as a second parameter



--------------------
handler interface
--------------------
The handler is nothing, just an interface
with only one method and that method has two parameters:
	1. response writer
	2. pointer of request.


type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}



-----------------------
Handle function
-----------------------
handle is a function with two parameters:
	1. A pattern of the route
	2. An object of a user-defined type that implements the handler interface

-----------------------
http.NewServeMux()
----------------------
http.NewServeMux() it just a function that returns :
	an object of a struct which implemented handle interface

*/

type person struct {
	name string
	age  int
}

// Implement handle interface
func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi all"))
}

// func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hi all"))
// }

const port = ":3885"

func main() {

	var p = &person{name: "Rajeev", age: 26}

	fmt.Println("server running on  http://localhost" + port)
	if err := http.ListenAndServe(port, p); err != nil {
		panic(err)
	}
}
