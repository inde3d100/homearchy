package cli

import (
	"github.com/spf13/cobra"
)

// DocsCmd is the CLI docs command for opening Homearchy documentation in the browser.
//
// macOS: uses open. Linux: uses xdg-open.
// Other platforms return CodeNotSupported.
var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open documentation in the browser",
	Long: `Open version-aware Homearchy documentation pages in the default browser.

The URL points to the documentation for the currently installed version
of Homearchy, so you always see the relevant reference.

Subcommands:
  cli      Open the CLI reference documentation
  config   Open the configuration reference documentation

Example:
  homearchy docs cli      Open CLI docs
  homearchy docs config   Open config docs`,
}

// DocsCLICmd is the CLI docs cli subcommand.
var DocsCLICmd = &cobra.Command{
	Use:   "cli",
	Short: "Open CLI documentation",
	Long:  "Open the Homearchy CLI documentation in the default browser.",
	RunE: func(_ *cobra.Command, _ []string) error {
		return openDocsPage("docs/CLI.md")
	},
}

// DocsConfigCmd is the CLI docs config subcommand.
var DocsConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"configuration"},
	Short:   "Open configuration documentation",
	Long:    "Open the Homearchy configuration documentation in the default browser.",
	RunE: func(_ *cobra.Command, _ []string) error {
		return openDocsPage("docs/CONFIGURATION.md")
	},
}

func init() {
	DocsCmd.AddCommand(DocsCLICmd)
	DocsCmd.AddCommand(DocsConfigCmd)
	RootCmd.AddCommand(DocsCmd)
}
