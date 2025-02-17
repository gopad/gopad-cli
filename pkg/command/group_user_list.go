package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type groupUserListBind struct {
	ID     string
	Format string
}

// tmplgroupUserList represents a row within group user listing.
var tmplgroupUserList = "Slug: \x1b[33m{{ .User.Username }} \x1b[0m" + `
ID: {{ .User.ID }}
Email: {{ .User.Email }}
Perm: {{ .Perm }}
`

var (
	groupUserListCmd = &cobra.Command{
		Use:   "list",
		Short: "List assigned users for a group",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, groupUserListAction)
		},
		Args: cobra.NoArgs,
	}

	groupUserListArgs = groupUserListBind{}
)

func init() {
	groupUserCmd.AddCommand(groupUserListCmd)

	groupUserListCmd.Flags().StringVarP(
		&groupUserListArgs.ID,
		"id",
		"i",
		"",
		"Group ID or slug",
	)

	groupUserListCmd.Flags().StringVar(
		&groupUserListArgs.Format,
		"format",
		tmplgroupUserList,
		"Custom output format",
	)
}

func groupUserListAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if groupUserListArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	resp, err := client.ListGroupUsersWithResponse(
		ccmd.Context(),
		groupUserListArgs.ID,
		&gopad.ListGroupUsersParams{
			Limit:  gopad.ToPtr(10000),
			Offset: gopad.ToPtr(0),
		},
	)

	if err != nil {
		return err
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		basicFuncMap,
	).Parse(
		fmt.Sprintln(groupUserListArgs.Format),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		records := resp.JSON200.Users

		if len(records) == 0 {
			fmt.Fprintln(os.Stderr, "Empty result")
			return nil
		}

		for _, record := range records {
			if err := tmpl.Execute(
				os.Stdout,
				record,
			); err != nil {
				return fmt.Errorf("failed to render template: %w", err)
			}
		}
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
