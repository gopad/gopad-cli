package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type userTeamAppendBind struct {
	ID   string
	Team string
	Perm string
}

var (
	userTeamAppendCmd = &cobra.Command{
		Use:   "append",
		Short: "Append team to user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userTeamAppendAction)
		},
		Args: cobra.NoArgs,
	}

	userTeamAppendArgs = userTeamAppendBind{}
)

func init() {
	userTeamCmd.AddCommand(userTeamAppendCmd)

	userTeamAppendCmd.Flags().StringVarP(
		&userTeamAppendArgs.ID,
		"id",
		"i",
		"",
		"User ID or slug",
	)

	userTeamAppendCmd.Flags().StringVar(
		&userTeamAppendArgs.Team,
		"team",
		"",
		"Team ID or slug",
	)

	userTeamAppendCmd.Flags().StringVar(
		&userTeamAppendArgs.Perm,
		"perm",
		"",
		"Role for the team",
	)
}

func userTeamAppendAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if userTeamAppendArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	if userTeamAppendArgs.Team == "" {
		return fmt.Errorf("you must provide a team ID or a slug")
	}

	body := gopad.AttachUserToTeamJSONRequestBody{
		Team: userTeamAppendArgs.Team,
	}

	if teamUserAppendArgs.Perm != "" {
		body.Perm = gopad.ToPtr(userTeamPerm(userTeamAppendArgs.Perm))
	}

	resp, err := client.AttachUserToTeamWithResponse(
		ccmd.Context(),
		userTeamAppendArgs.ID,
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
