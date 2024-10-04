package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type userTeamPermitBind struct {
	ID   string
	Team string
	Perm string
}

var (
	userTeamPermitCmd = &cobra.Command{
		Use:   "permit",
		Short: "Permit team for user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userTeamPermitAction)
		},
		Args: cobra.NoArgs,
	}

	userTeamPermitArgs = userTeamPermitBind{}
)

func init() {
	userTeamCmd.AddCommand(userTeamPermitCmd)

	userTeamPermitCmd.Flags().StringVarP(
		&userTeamPermitArgs.ID,
		"id",
		"i",
		"",
		"User ID or slug",
	)

	userTeamPermitCmd.Flags().StringVar(
		&userTeamPermitArgs.Team,
		"team",
		"",
		"Team ID or slug",
	)

	userTeamPermitCmd.Flags().StringVar(
		&userTeamPermitArgs.Perm,
		"perm",
		"",
		"Role for the team",
	)
}

func userTeamPermitAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if userTeamPermitArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	if userTeamPermitArgs.Team == "" {
		return fmt.Errorf("you must provide a team ID or a slug")
	}

	body := gopad.PermitUserTeamJSONRequestBody{
		Team: userTeamPermitArgs.Team,
	}

	if teamUserAppendArgs.Perm != "" {
		body.Perm = gopad.ToPtr(userTeamPerm(userTeamPermitArgs.Perm))
	}

	resp, err := client.PermitUserTeamWithResponse(
		ccmd.Context(),
		userTeamPermitArgs.ID,
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
