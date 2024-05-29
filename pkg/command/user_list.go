package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplUserList represents a row within user listing.
var tmplUserList = "Username: \x1b[33m{{ .Username }} \x1b[0m" + `
ID: {{ .Id }}
`

var (
	userListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all users",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userListAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userListCmd)

	userListCmd.Flags().String("format", tmplUserList, "Custom output format")
	_ = viper.BindPFlag("user.list.format", userListCmd.Flags().Lookup("format"))
}

func userListAction(ccmd *cobra.Command, _ []string, client *Client) error {
	resp, err := client.ListUsersWithResponse(
		ccmd.Context(),
		&gopad.ListUsersParams{
			Limit:  gopad.ToPtr(1000),
			Offset: gopad.ToPtr(0),
		},
	)

	if err != nil {
		return prettyError(err)
	}

	records := gopad.FromPtr(resp.JSON200.Users)

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
		fmt.Sprintln(viper.GetString("user.list.format")),
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
