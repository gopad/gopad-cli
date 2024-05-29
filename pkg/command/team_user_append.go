package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamUserAppendCmd = &cobra.Command{
		Use:   "append",
		Short: "Append user to team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamUserAppendAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamUserCmd.AddCommand(teamUserAppendCmd)

	teamUserAppendCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.user.append.id", teamUserAppendCmd.Flags().Lookup("id"))

	teamUserAppendCmd.Flags().StringP("user", "t", "", "User ID or slug")
	_ = viper.BindPFlag("team.user.append.user", teamUserAppendCmd.Flags().Lookup("user"))

	teamUserAppendCmd.Flags().String("perm", "", "Role for the user")
	_ = viper.BindPFlag("team.user.append.perm", teamUserAppendCmd.Flags().Lookup("perm"))
}

func teamUserAppendAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
