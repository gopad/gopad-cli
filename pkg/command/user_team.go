package command

import (
	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

var (
	userTeamCmd = &cobra.Command{
		Use:   "team",
		Short: "Team assignments",
		Args:  cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userTeamCmd)
}

func userTeamPerm(val string) gopad.UserTeamParamsPerm {
	switch val {
	case "owner":
		return gopad.UserTeamParamsPermOwner
	case "admin":
		return gopad.UserTeamParamsPermAdmin
	case "user":
		return gopad.UserTeamParamsPermUser
	}

	return gopad.UserTeamParamsPermUser
}
