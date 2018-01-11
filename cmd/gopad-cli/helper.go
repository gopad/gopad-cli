package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/gopad/gopad-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{
	"split":    strings.Split,
	"join":     strings.Join,
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"contains": strings.Contains,
	"replace":  strings.Replace,
	"teamList": func(s []*sdk.Team) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"userList": func(s []*sdk.User) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
}

// GetIdentifierParam checks and returns the record id/slug parameter.
func GetIdentifierParam(c *cli.Context) string {
	val := c.String("id")

	if val == "" {
		fmt.Println("Error: you must provide an id or a slug.")
		os.Exit(1)
	}

	return val
}

// GetUserParam checks and returns the user id/slug parameter.
func GetUserParam(c *cli.Context) string {
	val := c.String("user")

	if val == "" {
		fmt.Println("Error: you must provide a user id or slug.")
		os.Exit(1)
	}

	return val
}

// GetTeamParam checks and returns the team id/slug parameter.
func GetTeamParam(c *cli.Context) string {
	val := c.String("team")

	if val == "" {
		fmt.Println("Error: you must provide a team id or slug.")
		os.Exit(1)
	}

	return val
}

// GetPermParam checks and returns the permission parameter.
func GetPermParam(c *cli.Context) string {
	val := c.String("perm")

	if val == "" {
		fmt.Println("Error: you must provide a permission.")
		os.Exit(1)
	}

	for _, perm := range []string{"user", "admin", "owner"} {
		if perm == val {
			return val
		}
	}

	fmt.Println("Error: invalid permission, can be user, admin or owner.")
	os.Exit(1)

	return ""
}
