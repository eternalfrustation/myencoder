package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port,http.FileServer(http.Dir("/")))
}
