package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func newLicenseCmd() *cobra.Command {
	licenseCmd := &cobra.Command{
		Use:   "license",
		Short: "Manage LICENSE files",
	}
	licenseCmd.AddCommand(newLicenseCreateCmd())
	return licenseCmd
}

func newLicenseCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create <license-key>",
		Short: "Create a LICENSE file from a GitHub license template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newRESTClient()
			if err != nil {
				return err
			}
			return runLicenseCreate(client, args[0], gitConfigUserName)
		},
	}
}

type licenseResponse struct {
	Body string `json:"body"`
}

// runLicenseCreate fetches the license template identified by licenseKey
// and writes it to ./LICENSE, substituting [year] and [fullname]
// placeholders. userName supplies the [fullname] value.
func runLicenseCreate(client restGetter, licenseKey string, userName func() (string, error)) error {
	if _, err := os.Stat("LICENSE"); err == nil {
		return fmt.Errorf("LICENSE already exists")
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	var resp licenseResponse
	if err := client.Get("licenses/"+licenseKey, &resp); err != nil {
		return formatAPIError(err, "license", licenseKey)
	}

	fullName, err := userName()
	if err != nil {
		return err
	}

	body := renderLicenseBody(resp.Body, time.Now().Year(), fullName)

	return os.WriteFile("LICENSE", []byte(body), 0o644)
}

func renderLicenseBody(body string, year int, fullName string) string {
	body = strings.ReplaceAll(body, "[year]", strconv.Itoa(year))
	body = strings.ReplaceAll(body, "[fullname]", fullName)
	return body
}

func gitConfigUserName() (string, error) {
	out, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		return "", fmt.Errorf("git config user.name is not set")
	}
	name := strings.TrimSpace(string(out))
	if name == "" {
		return "", fmt.Errorf("git config user.name is not set")
	}
	return name, nil
}
