package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	vault "github.com/hashicorp/vault/api"
	"crypto/tls"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// command line options
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

	client, err := VaultClient(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	log.Print(client.Address())
}

func VaultClient(config Config) (*vault.Client, error) {
	var protocol string
	if config.TLS {
		protocol = "https"
	} else {
		protocol = "http"
	}

	transport := &http.Transport{}
	if !config.Insecure {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	// Generate a config for the Vault client
	vaultConfig := vault.Config{
		Address: fmt.Sprintf("%v://%v:%v", protocol, config.Server, config.Port),
		HttpClient: &http.Client{Transport: transport },
	}

	// initialize the client
	vaultClient, err := vault.NewClient(&vaultConfig)
	if err != nil {
		return vaultClient, nil
	}
	vaultClient.SetToken(config.Token)
	vaultClient.Auth()

	return vaultClient, nil
}
