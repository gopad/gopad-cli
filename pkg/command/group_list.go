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

// tmplgroupList represents a row within user listing.
var tmplgroupList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

type groupListBind struct {
	Format string
}

var (
	groupListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all groups",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, groupListAction)
		},
		Args: cobra.NoArgs,
	}

	groupListArgs = groupListBind{}
)

func init() {
	groupCmd.AddCommand(groupListCmd)

	groupListCmd.Flags().StringVar(
		&groupListArgs.Format,
		"format",
		tmplgroupList,
		"Custom output format",
	)
}

func groupListAction(ccmd *cobra.Command, _ []string, client *Client) error {
	resp, err := client.ListGroupsWithResponse(
		ccmd.Context(),
		&gopad.ListGroupsParams{
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
		fmt.Sprintln(groupListArgs.Format),
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
