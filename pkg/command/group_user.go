package command

import (
	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

var (
	groupUserCmd = &cobra.Command{
		Use:   "user",
		Short: "User assignments",
		Args:  cobra.NoArgs,
	}
)

func init() {
	groupCmd.AddCommand(groupUserCmd)
}

func groupUserPerm(val string) gopad.UserGroupPerm {
	res, err := gopad.ToUserGroupPerm(val)

	if err != nil {
		return gopad.UserGroupPermUser
	}

	return res
}
