package cmd

import "github.com/spf13/cobra"

// NewRootCmd builds the "init" root command with all subcommands
// registered.
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "init",
		Short:         "gh-init generates common project files (LICENSE, .gitignore, ...)",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(newLicenseCmd())
	root.AddCommand(newIgnoreCmd())
	return root
}
