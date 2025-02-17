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

type userGroupListBind struct {
	ID     string
	Format string
}

// tmplUserGroupList represents a row within user group listing.
var tmplUserGroupList = "Slug: \x1b[33m{{ .Group.Slug }} \x1b[0m" + `
ID: {{ .Group.ID }}
Name: {{ .Group.Name }}
Perm: {{ .Perm }}
`

var (
	userGroupListCmd = &cobra.Command{
		Use:   "list",
		Short: "List assigned groups for a user",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, userGroupListAction)
		},
		Args: cobra.NoArgs,
	}

	userGroupListArgs = userGroupListBind{}
)

func init() {
	userGroupCmd.AddCommand(userGroupListCmd)

	userGroupListCmd.Flags().StringVarP(
		&userGroupListArgs.ID,
		"id",
		"i",
		"",
		"User ID or slug",
	)

	userGroupListCmd.Flags().StringVar(
		&userGroupListArgs.Format,
		"format",
		tmplUserGroupList,
		"Custom output format",
	)
}

func userGroupListAction(ccmd *cobra.Command, _ []string, client *Client) error {
	if userGroupListArgs.ID == "" {
		return fmt.Errorf("you must provide an ID or a slug")
	}

	resp, err := client.ListUserGroupsWithResponse(
		ccmd.Context(),
		userGroupListArgs.ID,
		&gopad.ListUserGroupsParams{
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
		fmt.Sprintln(userGroupListArgs.Format),
	)

	if err != nil {
		return fmt.Errorf("failed to process template: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		records := resp.JSON200.Groups

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
