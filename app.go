package main

import (
	"log"
	"net/http"

	"github.com/albshin/overmatch/app/route"
)

func main() {
	r := route.Route()
	log.Println("Listening...")
	go http.ListenAndServe(":2000", http.RedirectHandler("https://localhost:3000/", 301))
	log.Fatal(http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", r))
}
