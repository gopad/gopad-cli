package command

import (
	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

var (
	teamUserCmd = &cobra.Command{
		Use:   "user",
		Short: "User assignments",
		Args:  cobra.NoArgs,
	}
)

func init() {
	teamCmd.AddCommand(teamUserCmd)
}

func teamUserPerm(val string) gopad.TeamUserParamsPerm {
	switch val {
	case "owner":
		return gopad.TeamUserParamsPermOwner
	case "admin":
		return gopad.TeamUserParamsPermAdmin
	case "user":
		return gopad.TeamUserParamsPermUser
	}

	return gopad.TeamUserParamsPermUser
}
