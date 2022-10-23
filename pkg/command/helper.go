package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/urfave/cli/v2"
)

// sprigFuncMap provides template helpers provided by sprig.
var sprigFuncMap = sprig.TxtFuncMap()

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{}

// GetIdentifierParam checks and returns the record id/slug parameter.
func GetIdentifierParam(c *cli.Context) string {
	val := c.String("id")

	if val == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide an ID or a slug.\n")
		os.Exit(1)
	}

	return val
}

// GetUserParam checks and returns the user id/slug parameter.
func GetUserParam(c *cli.Context) string {
	val := c.String("user")

	if val == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide a user ID or slug.\n")
		os.Exit(1)
	}

	return val
}

// GetTeamParam checks and returns the team id/slug parameter.
func GetTeamParam(c *cli.Context) string {
	val := c.String("team")

	if val == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide a team ID or slug.\n")
		os.Exit(1)
	}

	return val
}
