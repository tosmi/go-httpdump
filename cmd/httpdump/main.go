package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	fmt.Println("Starting listener...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	b, err := httputil.DumpRequest(request, true)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Println(string(b))
}
