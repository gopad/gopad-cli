package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userTeamRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove team from user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userTeamRemoveAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userTeamCmd.AddCommand(userTeamRemoveCmd)

	userTeamRemoveCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.team.remove.id", userTeamRemoveCmd.Flags().Lookup("id"))

	userTeamRemoveCmd.Flags().StringP("team", "t", "", "Team ID or slug")
	_ = viper.BindPFlag("user.team.remove.team", userTeamRemoveCmd.Flags().Lookup("team"))
}

func userTeamRemoveAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
