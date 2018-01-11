package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gopad/gopad-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// HandleFunc is the real handle implementation.
type HandleFunc func(c *cli.Context, client sdk.ClientAPI) error

// Handle wraps the command function handler.
func Handle(c *cli.Context, fn HandleFunc) error {
	var (
		server = c.String("server")
		token  = c.String("token")

		client sdk.ClientAPI
	)

	if server == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide the server address.\n")
		os.Exit(1)
	}

	if _, err := url.Parse(server); err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid server address, bad format?.\n")
		os.Exit(1)
	}

	if token == "" {
		client = sdk.NewClient(
			server,
		)
	} else {
		client = sdk.NewClientToken(
			server,
			token,
		)
	}

	if err := fn(c, client); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(2)
	}

	return nil
}
