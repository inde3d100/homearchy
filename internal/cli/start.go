package cli

// StartCmd is the CLI start command.
var StartCmd = BuildSimpleCommand(
	"start",
	"Resume Homearchy after being stopped",
	`Resume the Homearchy daemon after it was paused with 'homearchy stop'.

This re-enables all navigation modes and actions without restarting
the daemon process. Use 'homearchy stop' to pause.`,
	"start",
)

func init() {
	RootCmd.AddCommand(StartCmd)
}
