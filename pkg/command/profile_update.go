package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	profileUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update profile details",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, profileUpdateAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	profileCmd.AddCommand(profileUpdateCmd)

	profileUpdateCmd.Flags().String("username", "", "Username for your profile")
	_ = viper.BindPFlag("profile.update.username", profileUpdateCmd.Flags().Lookup("username"))

	profileUpdateCmd.Flags().String("password", "", "Password for your profile")
	_ = viper.BindPFlag("profile.update.password", profileUpdateCmd.Flags().Lookup("password"))

	profileUpdateCmd.Flags().String("email", "", "Email for your profile")
	_ = viper.BindPFlag("profile.update.email", profileUpdateCmd.Flags().Lookup("email"))

	profileUpdateCmd.Flags().String("fullanme", "", "Fullname for your profile")
	_ = viper.BindPFlag("profile.update.fullanme", profileUpdateCmd.Flags().Lookup("fullanme"))
}

func profileUpdateAction(_ *cobra.Command, _ []string, _ *Client) error {
	// resp, err := client.Profile.ShowProfile(
	// 	profile.NewShowProfileParams(),
	// 	client.AuthInfo,
	// )

	// if err != nil {
	// 	switch val := err.(type) {
	// 	case *profile.ShowProfileForbidden:
	// 		return fmt.Errorf(*val.Payload.Message)
	// 	case *profile.ShowProfileDefault:
	// 		return fmt.Errorf(*val.Payload.Message)
	// 	default:
	// 		return PrettyError(err)
	// 	}
	// }

	// record := resp.Payload
	// changed := false

	// if val := c.String("slug"); c.IsSet("slug") && val != *record.Slug {
	// 	record.Slug = &val
	// 	changed = true
	// }

	// if val := c.String("email"); c.IsSet("email") && val != *record.Email {
	// 	record.Email = &val
	// 	changed = true
	// }

	// if val := c.String("username"); c.IsSet("username") && val != *record.Username {
	// 	record.Username = &val
	// 	changed = true
	// }

	// if val := c.String("password"); c.IsSet("password") {
	// 	password := strfmt.Password(val)
	// 	record.Password = &password
	// 	changed = true
	// }

	// if changed {
	// 	if err := record.Validate(strfmt.Default); err != nil {
	// 		return ValidteError(err)
	// 	}

	// 	_, err := client.Profile.UpdateProfile(
	// 		profile.NewUpdateProfileParams().WithProfile(record),
	// 		client.AuthInfo,
	// 	)

	// 	if err != nil {
	// 		switch val := err.(type) {
	// 		case *profile.UpdateProfileForbidden:
	// 			return fmt.Errorf(*val.Payload.Message)
	// 		case *profile.UpdateProfileDefault:
	// 			return fmt.Errorf(*val.Payload.Message)
	// 		case *profile.UpdateProfileUnprocessableEntity:
	// 			return ValidteError(*val.Payload)
	// 		default:
	// 			return PrettyError(err)
	// 		}
	// 	}

	// 	fmt.Fprintln(os.Stderr, "successfully update")
	// } else {
	// 	fmt.Fprintln(os.Stderr, "nothing to update...")
	// }

	return nil
}
