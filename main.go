package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aitva/oauth_to_jwt/handler"
)

// OAuthConfig contains the OAuth2 credential for identification.
type OAuthConfig struct {
	ID       string
	Code     string
	Redirect string
}

// Config contains the app configuration.
type Config struct {
	Address string
	OAuth   OAuthConfig
}

// LoadConfig loads configuration from disk.
func LoadConfig(filename string) (conf Config, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		return
	}
	return
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalln("fail to read config.json:", err)
	}
	if redirect := os.Getenv("OAUTH_REDIRECT"); redirect != "" {
		config.OAuth.Redirect = redirect
	}

	oauth := handler.NewOAuth(config.OAuth.ID, config.OAuth.Code, config.OAuth.Redirect)
	url := oauth.GetURL()
	log.Println("OAuth URL:", url)
	http.Handle("/", oauth)

	log.Println("listening on", config.Address)
	log.Fatalln(http.ListenAndServe(config.Address, nil))
}
