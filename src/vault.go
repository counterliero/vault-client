package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	vault "github.com/hashicorp/vault/api"
)

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
