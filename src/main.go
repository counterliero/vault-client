package main

import "log"

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(config)
}