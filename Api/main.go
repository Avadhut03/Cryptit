package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome to homepage")
	fmt.Println("Endpoint hit: homepage")
}
func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:3000",nil)
}