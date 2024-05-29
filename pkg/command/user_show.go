package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplUserShow represents a user within details view.
var tmplUserShow = "Username: \x1b[33m{{ .Username }} \x1b[0m" + `
ID: {{ .Id }}
Email: {{ .Email }}
{{- with .Fullname }}
Fullname: {{ .Fullname }}
Active: {{ .Active }}
Admin: {{ .Admin }}
Created: {{ .CreatedAt.AsTime }}
Updated: {{ .UpdatedAt.AsTime }}
`

var (
	userShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show an user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userShowAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userShowCmd)

	userShowCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.show.id", userShowCmd.Flags().Lookup("id"))

	userShowCmd.Flags().String("format", tmplUserShow, "Custom output format")
	_ = viper.BindPFlag("user.show.format", userShowCmd.Flags().Lookup("format"))
}

func userShowAction(ccmd *cobra.Command, _ []string, client *Client) error {
	resp, err := client.ShowUserWithResponse(
		ccmd.Context(),
		GetIdentifierParam("user.show.id"),
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
		fmt.Sprintln(viper.GetString("user.show.format")),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, gopad.FromPtr(resp.JSON200)); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
