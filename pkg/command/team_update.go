package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamUpdateAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamCmd.AddCommand(teamUpdateCmd)

	teamUpdateCmd.Flags().StringP("id", "i", "", "Team ID or slug")
	_ = viper.BindPFlag("team.update.id", teamUpdateCmd.Flags().Lookup("id"))

	teamUpdateCmd.Flags().String("slug", "", "Slug for team")
	_ = viper.BindPFlag("team.update.slug", teamUpdateCmd.Flags().Lookup("slug"))

	teamUpdateCmd.Flags().String("name", "", "Name for team")
	_ = viper.BindPFlag("team.update.name", teamUpdateCmd.Flags().Lookup("name"))
}

func teamUpdateAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
