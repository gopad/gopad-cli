package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userCreateAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().String("username", "", "Username for user")
	_ = viper.BindPFlag("user.create.username", userCreateCmd.Flags().Lookup("username"))

	userCreateCmd.Flags().String("password", "", "Password for user")
	_ = viper.BindPFlag("user.create.password", userCreateCmd.Flags().Lookup("password"))

	userCreateCmd.Flags().String("email", "", "Email for user")
	_ = viper.BindPFlag("user.create.email", userCreateCmd.Flags().Lookup("email"))

	userCreateCmd.Flags().String("fullname", "", "Fullname for user")
	_ = viper.BindPFlag("user.create.fullname", userCreateCmd.Flags().Lookup("fullname"))

	userCreateCmd.Flags().Bool("active", false, "Mark user as active")
	_ = viper.BindPFlag("user.create.active", userCreateCmd.Flags().Lookup("active"))

	userCreateCmd.Flags().Bool("inactive", false, "Mark user as inactive")
	_ = viper.BindPFlag("user.create.inactive", userCreateCmd.Flags().Lookup("inactive"))

	userCreateCmd.Flags().Bool("admin", false, "Mark user as admin")
	_ = viper.BindPFlag("user.create.admin", userCreateCmd.Flags().Lookup("admin"))

	userCreateCmd.Flags().Bool("regular", false, "Mark user as regular")
	_ = viper.BindPFlag("user.create.regular", userCreateCmd.Flags().Lookup("regular"))
}

func userCreateAction(_ *cobra.Command, _ []string, _ *Client) error {
	return nil
}
