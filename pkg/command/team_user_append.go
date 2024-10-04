package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type teamUserAppendBind struct {
	ID   string
	User string
	Perm string
}

var (
	teamUserAppendCmd = &cobra.Command{
		Use:   "append",
		Short: "Append user to team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamUserAppendAction)
		},
		Args: cobra.NoArgs,
	}

	teamUserAppendArgs = teamUserAppendBind{}
)

func init() {
	teamUserCmd.AddCommand(teamUserAppendCmd)

	teamUserAppendCmd.Flags().StringVarP(
		&teamUserAppendArgs.ID,
		"id",
		"i",
		"",
		"Team ID or slug",
	)

	teamUserAppendCmd.Flags().StringVar(
		&teamUserAppendArgs.User,
		"user",
		"",
		"User ID or slug",
	)

	teamUserAppendCmd.Flags().StringVar(
		&teamUserAppendArgs.Perm,
		"perm",
		"",
		"Role for the user",
	)
}

func teamUserAppendAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if teamUserAppendArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	if teamUserAppendArgs.User == "" {
		return fmt.Errorf("you must provide a user ID or a slug")
	}

	body := gopad.AttachTeamToUserJSONRequestBody{
		User: teamUserAppendArgs.User,
	}

	if teamUserAppendArgs.Perm != "" {
		body.Perm = gopad.ToPtr(teamUserPerm(teamUserAppendArgs.Perm))
	}

	resp, err := client.AttachTeamToUserWithResponse(
		ccmd.Context(),
		teamUserAppendArgs.ID,
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
		return errors.New(gopad.FromPtr(resp.JSON403.Message))
	case http.StatusNotFound:
		return errors.New(gopad.FromPtr(resp.JSON404.Message))
	case http.StatusInternalServerError:
		return errors.New(gopad.FromPtr(resp.JSON500.Message))
	default:
		return fmt.Errorf("unknown api response")
	}

	return nil
}
