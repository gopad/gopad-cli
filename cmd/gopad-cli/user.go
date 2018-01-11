package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// userFuncMap provides template helper functions.
var userFuncMap = template.FuncMap{}

// tmplUserList represents a row within user listing.
var tmplUserList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
`

// tmplUserShow represents a user within details view.
var tmplUserShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
Email: {{ .Email }}
Active: {{ .Active }}
Admin: {{ .Admin }}{{with .Teams}}
Teams: {{ teamList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplUserTeamList represents a row within user team listing.
var tmplUserTeamList = "Slug: \x1b[33m{{ .User.Slug }} \x1b[0m" + `
ID: {{ .User.ID }}
Name: {{ .User.Name }}
Permission: {{ .Perm }}
`

// User provides the sub-command for the user API.
func User() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "User related sub-commands",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "list all users",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Value: tmplUserList,
						Usage: "custom output format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserList)
				},
			},
			{
				Name:      "show",
				Usage:     "show an user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "user id or slug",
					},
					&cli.StringFlag{
						Name:  "format",
						Value: tmplUserShow,
						Usage: "custom output format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserShow)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "delete an user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "user id or slug",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserDelete)
				},
			},
			{
				Name:      "update",
				Usage:     "update an user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "user id or slug",
					},
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "provide a slug",
					},
					&cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "provide an username",
					},
					&cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "provide an email",
					},
					&cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "provide a password",
					},
					&cli.BoolFlag{
						Name:  "active",
						Usage: "mark user as active",
					},
					&cli.BoolFlag{
						Name:  "blocked",
						Usage: "mark user as blocked",
					},
					&cli.BoolFlag{
						Name:  "admin",
						Usage: "mark user as admin",
					},
					&cli.BoolFlag{
						Name:  "user",
						Usage: "mark user as user",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserUpdate)
				},
			},
			{
				Name:      "create",
				Usage:     "create an user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "provide a slug",
					},
					&cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "provide an username",
					},
					&cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "provide an email",
					},
					&cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "provide a password",
					},
					&cli.BoolFlag{
						Name:  "active",
						Usage: "mark user as active",
					},
					&cli.BoolFlag{
						Name:  "blocked",
						Usage: "mark user as blocked",
					},
					&cli.BoolFlag{
						Name:  "admin",
						Usage: "mark user as admin",
					},
					&cli.BoolFlag{
						Name:  "user",
						Usage: "mark user as user",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserCreate)
				},
			},
			{
				Name:  "team",
				Usage: "team assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "list assigned teams for a user",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "user id or slug",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplUserTeamList,
								Usage: "custom output format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, UserTeamList)
						},
					},
					{
						Name:      "append",
						Usage:     "append a team to an user",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "user id or slug",
							},
							&cli.StringFlag{
								Name:  "team, t",
								Value: "",
								Usage: "team id or slug",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "permission, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, UserTeamAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "update user team permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "user id or slug to update",
							},
							&cli.StringFlag{
								Name:  "team, t",
								Value: "",
								Usage: "team id or slug to update",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "permission, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, UserTeamPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "remove a team from an user",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "user id or slug to remove from",
							},
							&cli.StringFlag{
								Name:  "team, t",
								Value: "",
								Usage: "team id or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, UserTeamRemove)
						},
					},
				},
			},
		},
	}
}

// UserList provides the sub-command to list all users.
func UserList(c *cli.Context, client sdk.ClientAPI) error {
	records, err := client.UserList()

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserShow provides the sub-command to show user details.
func UserShow(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// UserDelete provides the sub-command to delete a user.
func UserDelete(c *cli.Context, client sdk.ClientAPI) error {
	err := client.UserDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully delete\n")
	return nil
}

// UserUpdate provides the sub-command to update a user.
func UserUpdate(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

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

	if c.IsSet("active") && c.IsSet("blocked") {
		return fmt.Errorf("conflict, you can mark it only active or blocked")
	}

	if c.IsSet("active") {
		record.Active = true
		changed = true
	}

	if c.IsSet("blocked") {
		record.Active = false
		changed = true
	}

	if c.IsSet("admin") && c.IsSet("user") {
		return fmt.Errorf("conflict, you can mark it only admin or user")
	}

	if c.IsSet("admin") {
		record.Admin = true
		changed = true
	}

	if c.IsSet("user") {
		record.Admin = false
		changed = true
	}

	if changed {
		_, patch := client.UserPatch(
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

// UserCreate provides the sub-command to create a user.
func UserCreate(c *cli.Context, client sdk.ClientAPI) error {
	record := &sdk.User{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("username"); c.IsSet("username") && val != "" {
		record.Username = val
	} else {
		return fmt.Errorf("you must provide an username")
	}

	if val := c.String("email"); c.IsSet("email") && val != "" {
		record.Email = val
	} else {
		return fmt.Errorf("you must provide an email")
	}

	if val := c.String("password"); c.IsSet("password") && val != "" {
		record.Password = val
	} else {
		return fmt.Errorf("you must provide a password")
	}

	if c.IsSet("active") && c.IsSet("blocked") {
		return fmt.Errorf("conflict, you can mark it only active or blocked")
	}

	if c.IsSet("active") {
		record.Active = true
	}

	if c.IsSet("blocked") {
		record.Active = false
	}

	if c.IsSet("admin") && c.IsSet("user") {
		return fmt.Errorf("conflict, you can mark it only admin or user")
	}

	if c.IsSet("admin") {
		record.Admin = true
	}

	if c.IsSet("user") {
		record.Admin = false
	}

	_, err := client.UserPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully created\n")
	return nil
}

// UserTeamList provides the sub-command to list teams of the user.
func UserTeamList(c *cli.Context, client sdk.ClientAPI) error {
	records, err := client.UserTeamList(
		sdk.UserTeamParams{
			User: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserTeamAppend provides the sub-command to append a team to the user.
func UserTeamAppend(c *cli.Context, client sdk.ClientAPI) error {
	err := client.UserTeamAppend(
		sdk.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully appended to team\n")
	return nil
}

// UserTeamPerm provides the sub-command to update user team permissions.
func UserTeamPerm(c *cli.Context, client sdk.ClientAPI) error {
	err := client.UserTeamPerm(
		sdk.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully updated permissions\n")
	return nil
}

// UserTeamRemove provides the sub-command to remove a team from the user.
func UserTeamRemove(c *cli.Context, client sdk.ClientAPI) error {
	err := client.UserTeamDelete(
		sdk.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully removed from team\n")
	return nil
}
