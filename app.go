package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/albshin/overpug/app/controller"
	"github.com/albshin/overpug/app/route"
	"github.com/albshin/overpug/app/session"
)

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err := decoder.Decode(config)
	if err != nil {
		panic(err)
	}
	log.Println("Config loaded.")
	controller.ConfigureOAuth2(config.OAuth2)

	// Create router, middleware before routing
	r := route.Middleware(route.Route())
	log.Println("Listening...")
	go http.ListenAndServe(":2000", http.RedirectHandler("https://localhost:3000/", 301))
	log.Fatal(http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", r))
}

type Configuration struct {
	Session session.SessionInfo   `json:"Session"`
	OAuth2  controller.OAuth2Info `json:"OAuth2"`
}
