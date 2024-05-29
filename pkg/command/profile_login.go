package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplProfileLogin represents a expiring login token.
var tmplProfileLogin = "Token: \x1b[33m{{ .Token }} \x1b[0m" + `
Expires: {{ .ExpiresAt }}
`

var (
	profileLoginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login by credentials",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, profileLoginAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	profileCmd.AddCommand(profileLoginCmd)

	profileLoginCmd.Flags().String("username", "", "Username for authentication")
	_ = viper.BindPFlag("profile.login.username", profileLoginCmd.Flags().Lookup("username"))

	profileLoginCmd.Flags().String("password", "", "Password for authentication")
	_ = viper.BindPFlag("profile.login.password", profileLoginCmd.Flags().Lookup("password"))

	profileLoginCmd.Flags().String("format", tmplProfileLogin, "Custom output format")
	_ = viper.BindPFlag("profile.login.format", profileLoginCmd.Flags().Lookup("format"))
}

func profileLoginAction(_ *cobra.Command, _ []string, _ *Client) error {
	if !viper.IsSet("profile.login.username") {
		return fmt.Errorf("please provide a username")
	}

	if !viper.IsSet("profile.login.password") {
		return fmt.Errorf("please provide a password")
	}

	// resp, err := client.Auth.LoginUser(
	// 	auth.NewLoginUserParams().WithAuthLogin(&models.AuthLogin{
	// 		Username: viper.GetString("profile.login.username"),
	// 		Password: viper.GetString("profile.login.password"),
	// 	}),
	// )

	// if err != nil {
	// 	switch val := err.(type) {
	// 	case *auth.LoginUserUnauthorized:
	// 		return fmt.Errorf(*val.Payload.Message)
	// 	case *auth.LoginUserDefault:
	// 		return fmt.Errorf(*val.Payload.Message)
	// 	default:
	// 		return PrettyError(err)
	// 	}
	// }

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		basicFuncMap,
	).Parse(
		fmt.Sprintln(viper.GetString("profile.login.format")),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
