package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type userGroupAppendBind struct {
	ID    string
	Group string
	Perm  string
}

var (
	userGroupAppendCmd = &cobra.Command{
		Use:   "append",
		Short: "Append group to user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userGroupAppendAction)
		},
		Args: cobra.NoArgs,
	}

	userGroupAppendArgs = userGroupAppendBind{}
)

func init() {
	userGroupCmd.AddCommand(userGroupAppendCmd)

	userGroupAppendCmd.Flags().StringVarP(
		&userGroupAppendArgs.ID,
		"id",
		"i",
		"",
		"User ID or slug",
	)

	userGroupAppendCmd.Flags().StringVar(
		&userGroupAppendArgs.Group,
		"group",
		"",
		"Group ID or slug",
	)

	userGroupAppendCmd.Flags().StringVar(
		&userGroupAppendArgs.Perm,
		"perm",
		"",
		"Role for the group",
	)
}

func userGroupAppendAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if userGroupAppendArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	if userGroupAppendArgs.Group == "" {
		return fmt.Errorf("you must provide a group ID or a slug")
	}

	body := gopad.AttachUserToGroupJSONRequestBody{
		Group: userGroupAppendArgs.Group,
	}

	if groupUserAppendArgs.Perm != "" {
		body.Perm = string(userGroupPerm(userGroupAppendArgs.Perm))
	} else {
		body.Perm = string(gopad.UserGroupPermUser)
	}

	resp, err := client.AttachUserToGroupWithResponse(
		ccmd.Context(),
		userGroupAppendArgs.ID,
		body,
	)

	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		fmt.Fprintln(os.Stderr, gopad.FromPtr(resp.JSON200.Message))
	case http.StatusUnprocessableEntity:
		return validationError(resp.JSON422)
	case http.StatusPreconditionFailed:
		return errors.New(gopad.FromPtr(resp.JSON412.Message))
	case http.StatusForbidden:
		if resp.JSON403 != nil {
			return errors.New(gopad.FromPtr(resp.JSON403.Message))
		}

		return errors.New(http.StatusText(http.StatusForbidden))
	case http.StatusNotFound:
		if resp.JSON404 != nil {
			return errors.New(gopad.FromPtr(resp.JSON404.Message))
		}

		return errors.New(http.StatusText(http.StatusNotFound))
	case http.StatusInternalServerError:
		if resp.JSON500 != nil {
			return errors.New(gopad.FromPtr(resp.JSON500.Message))
		}

		return errors.New(http.StatusText(http.StatusInternalServerError))
	case http.StatusUnauthorized:
		return ErrMissingRequiredCredentials
	default:
		return ErrUnknownServerResponse
	}

	return nil
}
