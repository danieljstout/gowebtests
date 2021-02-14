package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Listening at port 8080. Try http://localhost:8080/dickbag")
	http.HandleFunc("/", serveHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
