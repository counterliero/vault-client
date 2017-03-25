package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// command line options
	var secretPath string
	var secretKey string
	flag.StringVar(&secretPath, "secret", "", "Path to secret")
	flag.StringVar(&secretKey, "key", "", "Key of the secret to get")
	flag.StringVar(&config.Server, "server", config.Server, "URL to Vault server")
	flag.IntVar(&config.Port, "port", config.Port, "Port that Vault listens on")
	flag.StringVar(&config.Token, "token", config.Token, "A valid Vault token")
	flag.BoolVar(&config.TLS, "tls", config.TLS, "Use SSL/TLS")
	flag.BoolVar(&config.Insecure, "insecure", config.Insecure, "Don't verify SSL/TLS connections")
	flag.Parse()

	if len(config.Token) == 0 {
		fmt.Fprintln(os.Stderr, "Missing mandatory parameter: token")
		os.Exit(1)
	}
	if len(secretPath) == 0 {
		fmt.Fprintln(os.Stderr, "Missing mandatory parameter: secret")
		os.Exit(1)
	}
	if len(secretKey) == 0 {
		fmt.Fprintln(os.Stderr, "Missing mandatory parameter: key")
		os.Exit(1)
	}

	client, err := VaultClient(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	secret, err := ReadSecret(*client, secretPath, secretKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error(), secretPath)
		os.Exit(1)
	}

	fmt.Println(secret)
}
