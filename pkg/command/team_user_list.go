package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplTeamUserList represents a row within team user listing.
var tmplTeamUserList = "Slug: \x1b[33m{{ .UserSlug }} \x1b[0m" + `
ID: {{ .UserId }}
Username: {{ .UserName }}
`

var (
	teamUserListCmd = &cobra.Command{
		Use:   "list",
		Short: "List assigned users for a team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamUserListAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamUserCmd.AddCommand(teamUserListCmd)

	teamUserListCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.user.list.id", teamUserListCmd.Flags().Lookup("id"))

	teamUserListCmd.Flags().String("format", tmplTeamUserList, "Custom output format")
	_ = viper.BindPFlag("team.user.list.format", teamUserListCmd.Flags().Lookup("format"))
}

func teamUserListAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
