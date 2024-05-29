package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userTeamAppendCmd = &cobra.Command{
		Use:   "append",
		Short: "Append team to user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userTeamAppendAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userTeamCmd.AddCommand(userTeamAppendCmd)

	userTeamAppendCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.team.append.id", userTeamAppendCmd.Flags().Lookup("id"))

	userTeamAppendCmd.Flags().StringP("team", "t", "", "Team ID or slug")
	_ = viper.BindPFlag("user.team.append.team", userTeamAppendCmd.Flags().Lookup("team"))

	userTeamAppendCmd.Flags().String("perm", "", "Role for the team")
	_ = viper.BindPFlag("user.team.append.perm", userTeamAppendCmd.Flags().Lookup("perm"))
}

func userTeamAppendAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
