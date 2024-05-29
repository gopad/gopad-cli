package command

import (
	"strings"

	"github.com/gopad/gopad-cli/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultServer = "http://localhost:8080"
)

var (
	rootCmd = &cobra.Command{
		Use:           "gopad-cli",
		Short:         "Etherpad for markdown with go",
		Version:       version.String,
		SilenceErrors: false,
		SilenceUsage:  true,

		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func init() {
	cobra.OnInitialize(setupConfig)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show the help")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version")

	rootCmd.PersistentFlags().StringP("server", "s", defaultServer, "API server")
	viper.SetDefault("server", "")
	_ = viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))

	rootCmd.PersistentFlags().StringP("token", "t", "", "API token")
	viper.SetDefault("token", "")
	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// Run parses the command line arguments and executes the program.
func Run() error {
	return rootCmd.Execute()
}

func setupConfig() {
	viper.SetEnvPrefix("gopad_")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
