package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	r "ascii-art-web/functions"
)

func main() {
	url := "http://localhost:8080"

	fmt.Printf("Starting server at port 8080... \n")
	http.HandleFunc("/", r.Indexhandle)
	http.HandleFunc("/download", r.Download)
	http.HandleFunc("/ascii-art", r.Asciiart)
	http.HandleFunc("/404", r.NotFoundHandler)
	//for open browser tab
	exec.Command("xdg-open", url).Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}
