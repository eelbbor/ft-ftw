package main

import (
	"fmt"
	"ftftw/toggles"
	"net/http"
)

func main() {
	fmt.Println("Feature Toggles For The Win!")
	http.HandleFunc("/", http.HandlerFunc(toggles.Handler))
	http.ListenAndServe(":8080", nil)
}
