package command

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmplProfileShow represents a profile within details view.
var tmplProfileShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .Id }}
Username: {{ .Username }}
Email: {{ .Email }}
Firstname: {{ .Firstname }}
Lastname: {{ .Lastname }}
Active: {{ .Active }}
Admin: {{ .Admin }}
Created: {{ .CreatedAt }}
Updated: {{ .UpdatedAt }}
`

var (
	profileShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show profile details",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, profileShowAction)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	profileCmd.AddCommand(profileShowCmd)

	profileShowCmd.Flags().String("format", tmplProfileShow, "Custom output format")
	_ = viper.BindPFlag("profile.show.format", profileShowCmd.Flags().Lookup("format"))
}

func profileShowAction(_ *cobra.Command, _ []string, _ *Client) error {
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

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		basicFuncMap,
	).Parse(
		fmt.Sprintln(viper.GetString("profile.show.format")),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
