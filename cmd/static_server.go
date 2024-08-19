package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func StartStaticServer(staticDir string, port int) {
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", http.StripPrefix("/", fs))
	log.Printf("Server starting on port %d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
