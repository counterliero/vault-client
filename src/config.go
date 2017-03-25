package main

import (
	"os/user"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
	Token  string `json:"token"`
	TLS    bool   `json:"tls"`
}

func LoadConfig() (Config, error) {
	// set defaults
	config := Config{
		Server: "localhost",
		Port: 8200,
		Token: "some-token",
		TLS: true,
	}

	// get current user
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return config, err
	}

	// read configuration from home dir
	configFile, err := ioutil.ReadFile(usr.HomeDir + "/.vaultrc")
	if err != nil {
		log.Print("Did not find %s/.vaultrc, using default config.", usr.HomeDir)
		return config, nil
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
