package main

import (
	"fmt"
	"net/http"

	"github.com/stephieandrade/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	// http.HandleFunc("/", func(writer http.ResponseWriter, response *http.Request){
	// 	n, err := fmt.Fprintf(writer, "Hello, world!")
	// 	if(err !=nil){
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	// })

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide", handlers.Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) //dont care about error
}