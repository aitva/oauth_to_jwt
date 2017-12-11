package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Address string
}

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
	log.Println("listening on", config.Address)
	log.Fatalln(http.ListenAndServe(config.Address, nil))
}
