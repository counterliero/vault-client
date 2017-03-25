package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	vault "github.com/hashicorp/vault/api"
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

func ReadSecret(vaultClient vault.Client, path, key string) (string, error) {
	res, err := vaultClient.Logical().Read(path)
	if res == nil || err != nil {
		return "", errors.New("Could not read secret")
	}

	secret := res.Data[key]
	if secret == nil {
		return "", err
	}

	return secret.(string), nil
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
