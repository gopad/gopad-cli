package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamCreateAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	teamCmd.AddCommand(teamCreateCmd)

	teamCreateCmd.Flags().String("slug", "", "Slug for team")
	_ = viper.BindPFlag("team.create.slug", teamCreateCmd.Flags().Lookup("slug"))

	teamCreateCmd.Flags().String("name", "", "Name for team")
	_ = viper.BindPFlag("team.create.name", teamCreateCmd.Flags().Lookup("name"))
}

func teamCreateAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
