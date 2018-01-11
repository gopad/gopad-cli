package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// teamFuncMap provides template helper functions.
var teamFuncMap = template.FuncMap{}

// tmplTeamList represents a row within user listing.
var tmplTeamList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplTeamShow represents a user within details view.
var tmplTeamShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}{{with .Users}}
Users: {{ userList . }}{{end}}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplTeamUserList represents a row within team user listing.
var tmplTeamUserList = "Slug: \x1b[33m{{ .User.Slug }} \x1b[0m" + `
ID: {{ .User.ID }}
Username: {{ .User.Username }}
Permission: {{ .Perm }}
`

// Team provides the sub-command for the team API.
func Team() *cli.Command {
	return &cli.Command{
		Name:  "team",
		Usage: "team commands",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "list all teams",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Value: tmplTeamList,
						Usage: "custom output format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamList)
				},
			},
			{
				Name:      "show",
				Usage:     "show a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "team id or slug",
					},
					&cli.StringFlag{
						Name:  "format",
						Value: tmplTeamShow,
						Usage: "custom output format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamShow)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "delete a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "team id or slug",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamDelete)
				},
			},
			{
				Name:      "update",
				Usage:     "update a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "team id or slug",
					},
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "provide a slug",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamUpdate)
				},
			},
			{
				Name:      "create",
				Usage:     "create a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "provide a slug",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamCreate)
				},
			},
			{
				Name:  "user",
				Usage: "user assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "list assigned users for a team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "team id or slug",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplTeamUserList,
								Usage: "custom output format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserList)
						},
					},
					{
						Name:      "append",
						Usage:     "append an user to a team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "team id or slug",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "user id or slug",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "permission, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "update team user permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "team id or slug",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "user id or slug",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "permission, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "remove an user from a team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "team id or slug",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "user id or slug",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserRemove)
						},
					},
				},
			},
		},
	}
}

// TeamList provides the sub-command to list all teams.
func TeamList(c *cli.Context, client sdk.ClientAPI) error {
	records, err := client.TeamList()

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
		teamFuncMap,
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

// TeamShow provides the sub-command to show team details.
func TeamShow(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.TeamGet(
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
		teamFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// TeamDelete provides the sub-command to delete a team.
func TeamDelete(c *cli.Context, client sdk.ClientAPI) error {
	err := client.TeamDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully delete\n")
	return nil
}

// TeamUpdate provides the sub-command to update a team.
func TeamUpdate(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.TeamGet(
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

	if val := c.String("name"); c.IsSet("name") && val != record.Name {
		record.Name = val
		changed = true
	}

	if changed {
		_, patch := client.TeamPatch(
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

// TeamCreate provides the sub-command to create a team.
func TeamCreate(c *cli.Context, client sdk.ClientAPI) error {
	record := &sdk.Team{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("you must provide a name")
	}

	_, err := client.TeamPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully created\n")
	return nil
}

// TeamUserList provides the sub-command to list users of the team.
func TeamUserList(c *cli.Context, client sdk.ClientAPI) error {
	records, err := client.TeamUserList(
		sdk.TeamUserParams{
			Team: GetIdentifierParam(c),
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
		teamFuncMap,
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

// TeamUserAppend provides the sub-command to append a user to the team.
func TeamUserAppend(c *cli.Context, client sdk.ClientAPI) error {
	err := client.TeamUserAppend(
		sdk.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully appended to user\n")
	return nil
}

// TeamUserPerm provides the sub-command to update team user permissions.
func TeamUserPerm(c *cli.Context, client sdk.ClientAPI) error {
	err := client.TeamUserPerm(
		sdk.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully updated permissions\n")
	return nil
}

// TeamUserRemove provides the sub-command to remove a user from the team.
func TeamUserRemove(c *cli.Context, client sdk.ClientAPI) error {
	err := client.TeamUserDelete(
		sdk.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "successfully removed from user\n")
	return nil
}
