package client

import (
	vault "github.com/hashicorp/vault/api"
	"log"
)

func ClientInit() *vault.Client {
	config := vault.DefaultConfig()
	err := config.ReadEnvironment()
	if err != nil {
		log.Fatalf("Unable to read environment config: %v", err)
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}
	return client
}
