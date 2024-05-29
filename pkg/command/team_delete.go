package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamDeleteAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamCmd.AddCommand(teamDeleteCmd)

	teamDeleteCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.delete.id", teamDeleteCmd.Flags().Lookup("id"))
}

func teamDeleteAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
