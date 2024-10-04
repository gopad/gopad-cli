package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type teamDeleteBind struct {
	ID string
}

var (
	teamDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an team",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, teamDeleteAction)
		},
		Args: cobra.NoArgs,
	}

	teamDeleteArgs = teamDeleteBind{}
)

func init() {
	teamCmd.AddCommand(teamDeleteCmd)

	teamDeleteCmd.Flags().StringVarP(
		&teamDeleteArgs.ID,
		"id",
		"i",
		"",
		"Team ID or slug",
	)
}

func teamDeleteAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if teamDeleteArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	resp, err := client.DeleteTeamWithResponse(
		ccmd.Context(),
		teamDeleteArgs.ID,
	)

	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		fmt.Fprintln(os.Stderr, "successfully delete")
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
