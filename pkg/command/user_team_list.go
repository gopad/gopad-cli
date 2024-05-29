package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplUserTeamList represents a row within user team listing.
var tmplUserTeamList = "Slug: \x1b[33m{{ .TeamSlug }} \x1b[0m" + `
ID: {{ .TeamId }}
Name: {{ .TeamName }}
`

var (
	userTeamListCmd = &cobra.Command{
		Use:   "list",
		Short: "List assigned teams for a user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userTeamListAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userTeamCmd.AddCommand(userTeamListCmd)

	userTeamListCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.team.list.id", userTeamListCmd.Flags().Lookup("id"))

	userTeamListCmd.Flags().String("format", tmplUserTeamList, "Custom output format")
	_ = viper.BindPFlag("user.team.list.format", userTeamListCmd.Flags().Lookup("format"))
}

func userTeamListAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
