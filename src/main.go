package main

import (
	"log"
	"flag"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// command line options
	flag.StringVar(&config.Server, "server", config.Server, "URL to Vault server")
	flag.IntVar(&config.Port, "port", config.Port, "Port that Vault listens on")
	flag.StringVar(&config.Token, "token", config.Token, "A valid Vault token")
	flag.BoolVar(&config.TLS, "tls", config.TLS, "Use SSL/TLS")
	flag.BoolVar(&config.Insecure, "insecure", config.Insecure, "Don't verify SSL/TLS connections")
	flag.Parse()
	if len(config.Token) == 0 {
		log.Fatal("Missing mandatory parameter: token")
		os.Exit(1)
	}

	log.Print(config)
}
