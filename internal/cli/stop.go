package cli

// StopCmd is the CLI stop command.
var StopCmd = BuildSimpleCommand(
	"stop",
	"Pause Homearchy (daemon stays running)",
	`Pause the Homearchy daemon. All navigation modes and actions are disabled,
but the daemon process remains running in the background.

Use 'homearchy start' to resume functionality.
Use 'homearchy status' to check whether Homearchy is active or paused.`,
	"stop",
)

func init() {
	RootCmd.AddCommand(StopCmd)
}
