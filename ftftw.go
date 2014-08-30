package main

import (
	"fmt"
	"net/http"
	"ft-ftw/toggles"
)

func main() {
    fmt.Println("Feature Toggles For The Win!")
	http.HandleFunc("/", http.HandlerFunc(toggles.Handler))
	http.ListenAndServe(":8080", nil)
}
