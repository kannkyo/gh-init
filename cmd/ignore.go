package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newIgnoreCmd() *cobra.Command {
	ignoreCmd := &cobra.Command{
		Use:   "ignore",
		Short: "Work with GitHub .gitignore templates",
	}
	ignoreCmd.AddCommand(newIgnoreListCmd())
	ignoreCmd.AddCommand(newIgnoreViewCmd())
	ignoreCmd.AddCommand(newIgnoreCreateCmd())
	return ignoreCmd
}

func newIgnoreListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List available .gitignore template names",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newRESTClient()
			if err != nil {
				return err
			}
			names, err := fetchIgnoreTemplateNames(client)
			if err != nil {
				return err
			}
			for _, name := range names {
				fmt.Fprintln(cmd.OutOrStdout(), name)
			}
			return nil
		},
	}
}

func newIgnoreViewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "view <template>",
		Short: "Print the contents of a .gitignore template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newRESTClient()
			if err != nil {
				return err
			}
			source, err := fetchIgnoreTemplateSource(client, args[0])
			if err != nil {
				return err
			}
			fmt.Fprint(cmd.OutOrStdout(), source)
			return nil
		},
	}
}

func newIgnoreCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create <template>",
		Short: "Create a .gitignore file from a GitHub template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newRESTClient()
			if err != nil {
				return err
			}
			return runIgnoreCreate(client, args[0])
		},
	}
}

type ignoreTemplateResponse struct {
	Source string `json:"source"`
}

func fetchIgnoreTemplateNames(client restGetter) ([]string, error) {
	var names []string
	if err := client.Get("gitignore/templates", &names); err != nil {
		return nil, err
	}
	return names, nil
}

func fetchIgnoreTemplateSource(client restGetter, template string) (string, error) {
	var resp ignoreTemplateResponse
	if err := client.Get("gitignore/templates/"+template, &resp); err != nil {
		return "", formatAPIError(err, "template", template)
	}
	return resp.Source, nil
}

// runIgnoreCreate fetches the .gitignore template identified by template
// and writes it to ./.gitignore.
func runIgnoreCreate(client restGetter, template string) error {
	if _, err := os.Stat(".gitignore"); err == nil {
		return fmt.Errorf(".gitignore already exists")
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	source, err := fetchIgnoreTemplateSource(client, template)
	if err != nil {
		return err
	}

	return os.WriteFile(".gitignore", []byte(source), 0o644)
}
