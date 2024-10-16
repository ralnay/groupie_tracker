package main

import (
	"fmt"
	"log"
	"net/http"
	piscine "piscine/functions"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", piscine.ServeIndex)
	http.HandleFunc("/artist/", piscine.ServeIndex)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

