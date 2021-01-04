package main

import (
	"fmt"
	"go-k8s/messages"
	"net/http"
)

func main() {
	fmt.Println("Server running at port 8000")
	http.HandleFunc("/", handlerRoot)
	http.ListenAndServe(":8000", nil)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, messages.Greeting("Code.education Rocks!"))
}
