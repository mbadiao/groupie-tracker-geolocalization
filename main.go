package main

import (
	"fmt"
	handlers "groupie/Func-Handlers"
	"net/http"
)

const port = ":8081"

func main() {

	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.HandleFunc("/artist/", handlers.ArtistDetails)
	http.HandleFunc("/search/", handlers.Search)
	http.HandleFunc("/", handlers.Home)
	fmt.Println("(http://localhost:8081) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
