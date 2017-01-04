package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/albshin/overpugs/app/controller"
	"github.com/albshin/overpugs/app/route"
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

	r := route.Route()
	log.Println("Listening...")
	go http.ListenAndServe(":2000", http.RedirectHandler("https://localhost:3000/", 301))
	log.Fatal(http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", r))
}

type Configuration struct {
	OAuth2 controller.OAuth2Info `json:"OAuth2"`
}
