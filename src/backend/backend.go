package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func Run(addr string) {
	http.HandleFunc("/", helloWorld)
	fmt.Println("server started and listening to port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
