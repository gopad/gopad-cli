package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userTeamPermitCmd = &cobra.Command{
		Use:   "permit",
		Short: "Permit team for user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userTeamPermitAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userTeamCmd.AddCommand(userTeamPermitCmd)

	userTeamPermitCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.team.permit.id", userTeamPermitCmd.Flags().Lookup("id"))

	userTeamPermitCmd.Flags().StringP("team", "t", "", "Team ID or slug")
	_ = viper.BindPFlag("user.team.permit.team", userTeamPermitCmd.Flags().Lookup("team"))

	userTeamPermitCmd.Flags().String("perm", "", "Role for the team")
	_ = viper.BindPFlag("user.team.permit.perm", userTeamPermitCmd.Flags().Lookup("perm"))
}

func userTeamPermitAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
