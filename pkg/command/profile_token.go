package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplProfileToken represents a permanent login token.
var tmplProfileToken = "Token: \x1b[33m{{ .Token }} \x1b[0m" + `
`

var (
	profileTokenCmd = &cobra.Command{
		Use:   "token",
		Short: "Show your token",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, profileTokenAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	profileCmd.AddCommand(profileTokenCmd)

	profileTokenCmd.Flags().String("format", tmplProfileToken, "Custom output format")
	_ = viper.BindPFlag("profile.token.format", profileTokenCmd.Flags().Lookup("format"))
}

func profileTokenAction(_ *cobra.Command, _ []string, _ *Client) error {
	// resp, err := client.Profile.TokenProfile(
	// 	profile.NewTokenProfileParams(),
	// 	client.AuthInfo,
	// )

	// if err != nil {
	// 	switch val := err.(type) {
	// 	case *profile.TokenProfileForbidden:
	// 		return fmt.Errorf(*val.Payload.Message)
	// 	case *profile.TokenProfileInternalServerError:
	// 		return fmt.Errorf(*val.Payload.Message)
	// 	case *profile.TokenProfileDefault:
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
		fmt.Sprintln(viper.GetString("profile.token.format")),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
