package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplTeamList represents a row within user listing.
var tmplTeamList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .Id }}
Name: {{ .Name }}
`

var (
	teamListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all teams",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamListAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamCmd.AddCommand(teamListCmd)

	teamListCmd.Flags().String("format", tmplTeamList, "Custom output format")
	_ = viper.BindPFlag("team.list.format", teamListCmd.Flags().Lookup("format"))
}

func teamListAction(ccmd *cobra.Command, _ []string, client *Client) error {
	resp, err := client.ListTeamsWithResponse(
		ccmd.Context(),
		&gopad.ListTeamsParams{
			Limit:  gopad.ToPtr(1000),
			Offset: gopad.ToPtr(0),
		},
	)

	if err != nil {
		return prettyError(err)
	}

	records := gopad.FromPtr(resp.JSON200.Teams)

	if len(records) == 0 {
		fmt.Fprintln(os.Stderr, "Empty result")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		basicFuncMap,
	).Parse(
		fmt.Sprintln(viper.GetString("team.list.format")),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	for _, record := range records {
		if err := tmpl.Execute(os.Stdout, record); err != nil {
			return fmt.Errorf("failed to render template: %w", err)
		}
	}

	return nil
}
