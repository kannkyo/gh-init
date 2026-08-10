package cmd

import "github.com/spf13/cobra"

// NewRootCmd builds the "scaffolder" root command with all subcommands
// registered.
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "scaffolder",
		Short:         "gh-scaffolder generates common project files (LICENSE, .gitignore, ...)",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(newLicenseCmd())
	root.AddCommand(newIgnoreCmd())
	return root
}
