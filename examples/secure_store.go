package main

import (
	"fmt"
	"os"

	"github.com/openshift-online/ocm-sdk-go/authentication"
	"github.com/openshift-online/ocm-sdk-go/authentication/securestore"
)

func main() {
	// Create a context:
	clientId := "ocm-cli"

	// Create the connection, and remember to close it:
	token, err := authentication.InitiateAuthCode(clientId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get token: %v\n", err)
		os.Exit(1)
	}

	// Get a list of available backends on the current OS
	available := securestore.AvailableBackends()
	fmt.Printf("Available backends: %v\n", available)

	// Create bytes
	config := []byte("mybytestringagain")

	// Upsert to keyring
	securestore.UpsertConfigToKeyring(config)

	// Upsert again to keyring
	config = []byte(token)
	securestore.UpsertConfigToKeyring(config)

	// Read bytes back from Keyring
	readVal, _ := securestore.GetConfigFromKeyring()
	// Should be a token
	fmt.Printf("Read from keyring: %s\n", string(readVal))
}
