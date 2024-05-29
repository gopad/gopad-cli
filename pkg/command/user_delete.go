package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userDeleteAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userDeleteCmd)

	userDeleteCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.delete.id", userDeleteCmd.Flags().Lookup("id"))
}

func userDeleteAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
