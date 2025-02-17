package command

import (
	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

var (
	userGroupCmd = &cobra.Command{
		Use:   "group",
		Short: "Group assignments",
		Args:  cobra.NoArgs,
	}
)

func init() {
	userCmd.AddCommand(userGroupCmd)
}

func userGroupPerm(val string) gopad.UserGroupPerm {
	res, err := gopad.ToUserGroupPerm(val)

	if err != nil {
		return gopad.UserGroupPermUser
	}

	return res
}
