package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type userGroupPermitBind struct {
	ID    string
	Group string
	Perm  string
}

var (
	userGroupPermitCmd = &cobra.Command{
		Use:   "permit",
		Short: "Permit group for user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userGroupPermitAction)
		},
		Args: cobra.NoArgs,
	}

	userGroupPermitArgs = userGroupPermitBind{}
)

func init() {
	userGroupCmd.AddCommand(userGroupPermitCmd)

	userGroupPermitCmd.Flags().StringVarP(
		&userGroupPermitArgs.ID,
		"id",
		"i",
		"",
		"User ID or slug",
	)

	userGroupPermitCmd.Flags().StringVar(
		&userGroupPermitArgs.Group,
		"group",
		"",
		"Group ID or slug",
	)

	userGroupPermitCmd.Flags().StringVar(
		&userGroupPermitArgs.Perm,
		"perm",
		"",
		"Role for the group",
	)
}

func userGroupPermitAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if userGroupPermitArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	if userGroupPermitArgs.Group == "" {
		return fmt.Errorf("you must provide a group ID or a slug")
	}

	if userGroupPermitArgs.Perm == "" {
		return fmt.Errorf("you must provide a a permission level like user, admin or owner")
	}

	body := gopad.PermitUserGroupJSONRequestBody{
		Group: userGroupPermitArgs.Group,
		Perm:  string(userGroupPerm(userGroupPermitArgs.Perm)),
	}

	resp, err := client.PermitUserGroupWithResponse(
		ccmd.Context(),
		userGroupPermitArgs.ID,
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
