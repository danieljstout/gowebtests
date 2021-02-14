package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func logRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "These are the Headers:\n")

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	fmt.Fprintf(w, "Here's the body:\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Fprintf(w, newStr)
}

func main() {
	fmt.Println("Listening at port 8080. Try http://localhost:8080/dickbag")
	http.HandleFunc("/", logRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
