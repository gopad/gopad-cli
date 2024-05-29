package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplTeamShow represents a user within details view.
var tmplTeamShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .Id }}
Name: {{ .Name }}
Created: {{ .CreatedAt.AsTime }}
Updated: {{ .UpdatedAt.AsTime }}
`

var (
	teamShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show an team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamShowAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamCmd.AddCommand(teamShowCmd)

	teamShowCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.show.id", teamShowCmd.Flags().Lookup("id"))

	teamShowCmd.Flags().String("format", tmplTeamShow, "Custom output format")
	_ = viper.BindPFlag("team.show.format", teamShowCmd.Flags().Lookup("format"))
}

func teamShowAction(ccmd *cobra.Command, _ []string, client *Client) error {
	resp, err := client.ShowTeamWithResponse(
		ccmd.Context(),
		GetIdentifierParam("team.show.id"),
	)

	if err != nil {
		return prettyError(err)
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		basicFuncMap,
	).Parse(
		fmt.Sprintln(viper.GetString("team.show.format")),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, gopad.FromPtr(resp.JSON200)); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
