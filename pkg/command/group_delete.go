package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type groupDeleteBind struct {
	ID string
}

var (
	groupDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an group",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, groupDeleteAction)
		},
		Args: cobra.NoArgs,
	}

	groupDeleteArgs = groupDeleteBind{}
)

func init() {
	groupCmd.AddCommand(groupDeleteCmd)

	groupDeleteCmd.Flags().StringVarP(
		&groupDeleteArgs.ID,
		"id",
		"i",
		"",
		"Group ID or slug",
	)
}

func groupDeleteAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if groupDeleteArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	resp, err := client.DeleteGroupWithResponse(
		ccmd.Context(),
		groupDeleteArgs.ID,
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
