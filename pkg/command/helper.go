package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/drone/funcmap"
	"github.com/spf13/viper"
)

// basicFuncMap provides template helpers provided by library.
var basicFuncMap = funcmap.Funcs

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{}

// GetIdentifierParam checks and returns the record id/slug parameter.
func GetIdentifierParam(name string) string {
	val := viper.GetString(name)

	if val == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide an ID or a slug.\n")
		os.Exit(1)
	}

	return val
}

// GetUserParam checks and returns the user id/slug parameter.
func GetUserParam(name string) string {
	val := viper.GetString(name)

	if val == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide a user ID or a slug.\n")
		os.Exit(1)
	}

	return val
}

// GetTeamParam checks and returns the team id/slug parameter.
func GetTeamParam(name string) string {
	val := viper.GetString(name)

	if val == "" {
		fmt.Fprintf(os.Stderr, "Error: you must provide a team ID or a slug.\n")
		os.Exit(1)
	}

	return val
}
