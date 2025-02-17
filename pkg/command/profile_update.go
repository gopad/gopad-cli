package command

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gopad/gopad-go/gopad"
	"github.com/spf13/cobra"
)

type profileUpdateBind struct {
	Username string
	Password string
	Email    string
	Fullname string
}

var (
	profileUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update profile details",
		Run: func(ccmd *cobra.Command, args []string) {
			Handle(ccmd, args, profileUpdateAction)
		},
		Args: cobra.NoArgs,
	}

	profileUpdateArgs = profileUpdateBind{}
)

func init() {
	profileCmd.AddCommand(profileUpdateCmd)

	profileUpdateCmd.Flags().StringVar(
		&profileUpdateArgs.Username,
		"username",
		"",
		"Username for your profile",
	)

	profileUpdateCmd.Flags().StringVar(
		&profileUpdateArgs.Password,
		"password",
		"",
		"Password for your profile",
	)

	profileUpdateCmd.Flags().StringVar(
		&profileUpdateArgs.Email,
		"email",
		"",
		"Email for your profile",
	)

	profileUpdateCmd.Flags().StringVar(
		&profileUpdateArgs.Fullname,
		"fullname",
		"",
		"Fullname for your profile",
	)
}

func profileUpdateAction(ccmd *cobra.Command, _ []string, client *Client) error {
	body := gopad.UpdateProfileJSONRequestBody{}
	changed := false

	if val := profileUpdateArgs.Username; val != "" {
		body.Username = gopad.ToPtr(val)
		changed = true
	}

	if val := profileUpdateArgs.Password; val != "" {
		body.Password = gopad.ToPtr(val)
		changed = true
	}

	if val := profileUpdateArgs.Email; val != "" {
		body.Email = gopad.ToPtr(val)
		changed = true
	}

	if val := profileUpdateArgs.Fullname; val != "" {
		body.Fullname = gopad.ToPtr(val)
		changed = true
	}

	if !changed {
		fmt.Fprintln(os.Stderr, "Nothing to update...")
		return nil
	}

	resp, err := client.UpdateProfileWithResponse(
		ccmd.Context(),
		body,
	)

	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		fmt.Fprintln(os.Stderr, "Successfully updated")
	case http.StatusUnprocessableEntity:
		return validationError(resp.JSON422)
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
