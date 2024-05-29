package command

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"syscall"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Client simply wraps the openapi client including authentication.
type Client struct {
	*gopad.ClientWithResponses
}

// HandleFunc is the real handle implementation.
type HandleFunc func(ccmd *cobra.Command, args []string, client *Client) error

// Handle wraps the command function handler.
func Handle(ccmd *cobra.Command, args []string, fn HandleFunc) {
	if viper.GetString("server") == "" {
		fmt.Fprintf(os.Stderr, "Error: You must provide the server address.\n")
		os.Exit(1)
	}

	server, err := url.Parse(viper.GetString("server"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid server address, bad format?\n")
		os.Exit(1)
	}

	child, err := gopad.NewClientWithResponses(
		server.String(),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to initialize client library\n")
		os.Exit(1)
	}

	client := &Client{
		ClientWithResponses: child,
	}

	if err := fn(ccmd, args, client); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(2)
	}
}

func prettyError(err error) error {
	if val, ok := err.(net.Error); ok && val.Timeout() {
		return fmt.Errorf("connection to server timed out")
	}

	switch val := err.(type) {
	case *net.OpError:
		switch val.Op {
		case "dial":
			return fmt.Errorf("unknown host for server connection")
		case "read":
			return fmt.Errorf("connection to server had been refused")
		default:
			return fmt.Errorf("failed to connect to the server")
		}
	case syscall.Errno:
		switch val {
		case syscall.ECONNREFUSED:
			return fmt.Errorf("connection to server had been refused")
		default:
			return fmt.Errorf("failed to connect to the server")
		}
	case net.Error:
		return fmt.Errorf("failed to connect to the server")
	default:
		return err
	}
}
