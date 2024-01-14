package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"johnbakhmat.tech/templs"
)

func main () {
	port := 3000

	component := templs.Home()
	http.Handle("/home", templ.Handler(component))
    http.Handle("/", http.FileServer(http.Dir("styles")))
	fmt.Printf("Server is running on :%d", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        panic(err)
    }
}
