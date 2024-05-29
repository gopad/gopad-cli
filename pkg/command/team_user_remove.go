package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamUserRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove user from team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamUserRemoveAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamUserCmd.AddCommand(teamUserRemoveCmd)

	teamUserRemoveCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.user.remove.id", teamUserRemoveCmd.Flags().Lookup("id"))

	teamUserRemoveCmd.Flags().StringP("user", "t", "", "User ID or slug")
	_ = viper.BindPFlag("team.user.remove.user", teamUserRemoveCmd.Flags().Lookup("user"))
}

func teamUserRemoveAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
