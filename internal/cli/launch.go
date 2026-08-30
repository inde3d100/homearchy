package cli

import (
	"github.com/spf13/cobra"
)

// LaunchCmd is the CLI launch command.
var LaunchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Start the Homearchy daemon",
	Long: `Start the Homearchy daemon process.

This initializes the accessibility engine, overlay, and IPC server
to handle navigation modes and commands. Once launched, Homearchy runs
in the background until quit from the system tray menu.

Use 'homearchy stop' to pause functionality (daemon stays running).`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		launchProgram(cmd, configPath)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(LaunchCmd)
}
