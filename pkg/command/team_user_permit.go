package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamUserPermitCmd = &cobra.Command{
		Use:   "permit",
		Short: "Permit user for team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamUserPermitAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamUserCmd.AddCommand(teamUserPermitCmd)

	teamUserPermitCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.user.permit.id", teamUserPermitCmd.Flags().Lookup("id"))

	teamUserPermitCmd.Flags().StringP("user", "t", "", "User ID or slug")
	_ = viper.BindPFlag("team.user.permit.user", teamUserPermitCmd.Flags().Lookup("user"))

	teamUserPermitCmd.Flags().String("perm", "", "Role for the user")
	_ = viper.BindPFlag("team.user.permit.perm", teamUserPermitCmd.Flags().Lookup("perm"))
}

func teamUserPermitAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
