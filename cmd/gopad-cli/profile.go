package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// profileFuncMap provides template helper functions.
var profileFuncMap = template.FuncMap{}

// tmplProfileShow represents a profile within details view.
var tmplProfileShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
Email: {{ .Email }}
Active: {{ .Active }}
Admin: {{ .Admin }}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// Profile provides the sub-command for the profile API.
func Profile() *cli.Command {
	return &cli.Command{
		Name:  "profile",
		Usage: "profile commands",
		Subcommands: []*cli.Command{
			{
				Name:  "show",
				Usage: "show profile details",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Value: tmplProfileShow,
						Usage: "custom output format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ProfileShow)
				},
			},
			{
				Name:  "token",
				Usage: "show your token",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "username for authentication",
					},
					&cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "password for authentication",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ProfileToken)
				},
			},
			{
				Name:  "update",
				Usage: "update profile details",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "provide a slug",
					},
					&cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "provide a username",
					},
					&cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "provide a email",
					},
					&cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "provide a password",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ProfileUpdate)
				},
			},
		},
	}
}

// ProfileShow provides the sub-command to show profile details.
func ProfileShow(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.ProfileGet()

	if err != nil {
		return err
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		profileFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// ProfileToken provides the sub-command to show your token.
func ProfileToken(c *cli.Context, client sdk.ClientAPI) error {
	if !client.IsAuthenticated() {
		if !c.IsSet("username") {
			return fmt.Errorf("please provide a username")
		}

		if !c.IsSet("password") {
			return fmt.Errorf("please provide a password")
		}

		login, err := client.AuthLogin(
			c.String("username"),
			c.String("password"),
		)

		if err != nil {
			return err
		}

		client = sdk.NewClientToken(
			c.String("server"),
			login.Token,
		)
	}

	record, err := client.ProfileToken()

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s\n", record.Token)
	return nil
}

// ProfileUpdate provides the sub-command to update the profile.
func ProfileUpdate(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.ProfileGet()

	if err != nil {
		return err
	}

	changed := false

	if val := c.String("slug"); c.IsSet("slug") && val != record.Slug {
		record.Slug = val
		changed = true
	}

	if val := c.String("username"); c.IsSet("username") && val != record.Username {
		record.Username = val
		changed = true
	}

	if val := c.String("email"); c.IsSet("email") && val != record.Email {
		record.Email = val
		changed = true
	}

	if val := c.String("password"); c.IsSet("password") {
		record.Password = val
		changed = true
	}

	if changed {
		_, patch := client.ProfilePatch(
			record,
		)

		if patch != nil {
			return patch
		}

		fmt.Fprintf(os.Stderr, "successfully updated\n")
	} else {
		fmt.Fprintf(os.Stderr, "nothing to update...\n")
	}

	return nil
}
