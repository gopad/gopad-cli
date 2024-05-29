package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userUpdateAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringP("id", "i", "", "User ID or slug")
	_ = viper.BindPFlag("user.update.id", userUpdateCmd.Flags().Lookup("id"))

	userUpdateCmd.Flags().String("username", "", "Username for user")
	_ = viper.BindPFlag("user.update.username", userUpdateCmd.Flags().Lookup("username"))

	userUpdateCmd.Flags().String("password", "", "Password for user")
	_ = viper.BindPFlag("user.update.password", userUpdateCmd.Flags().Lookup("password"))

	userUpdateCmd.Flags().String("email", "", "Email for user")
	_ = viper.BindPFlag("user.update.email", userUpdateCmd.Flags().Lookup("email"))

	userUpdateCmd.Flags().String("fullname", "", "Fullname for user")
	_ = viper.BindPFlag("user.update.fullname", userUpdateCmd.Flags().Lookup("fullname"))

	userUpdateCmd.Flags().Bool("active", false, "Mark user as active")
	_ = viper.BindPFlag("user.update.active", userUpdateCmd.Flags().Lookup("active"))

	userUpdateCmd.Flags().Bool("inactive", false, "Mark user as inactive")
	_ = viper.BindPFlag("user.update.inactive", userUpdateCmd.Flags().Lookup("inactive"))

	userUpdateCmd.Flags().Bool("admin", false, "Mark user as admin")
	_ = viper.BindPFlag("user.update.admin", userUpdateCmd.Flags().Lookup("admin"))

	userUpdateCmd.Flags().Bool("regular", false, "Mark user as regular")
	_ = viper.BindPFlag("user.update.regular", userUpdateCmd.Flags().Lookup("regular"))
}

func userUpdateAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
