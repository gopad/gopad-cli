package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type userDeleteBind struct {
	ID string
}

var (
	userDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userDeleteAction)
		},
		Args: cobra.NoArgs,
	}

	userDeleteArgs = userDeleteBind{}
)

func init() {
	userCmd.AddCommand(userDeleteCmd)

	userDeleteCmd.Flags().StringVarP(
		&userDeleteArgs.ID,
		"id",
		"i",
		"",
		"User ID or slug",
	)
}

func userDeleteAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if userDeleteArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	resp, err := client.DeleteUserWithResponse(
		ccmd.Context(),
		userDeleteArgs.ID,
	)

	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		fmt.Fprintln(os.Stderr, "Successfully deleted")
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
