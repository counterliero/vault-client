package main

import (
	"os/user"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Token    string `json:"token"`
	TLS      bool   `json:"tls"`
	Insecure bool   `json:"insecure"`
}

func LoadConfig() (Config, error) {
	// set defaults
	config := Config{
		Server: "localhost",
		Port: 8200,
		Insecure: true,
	}

	// get current user
	usr, err := user.Current()
	if err != nil {
		return config, err
	}

	// read configuration from home dir
	configFile, err := ioutil.ReadFile(usr.HomeDir + "/.vaultrc")
	if err != nil {
		return config, nil
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
