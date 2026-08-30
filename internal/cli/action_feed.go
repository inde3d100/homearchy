package cli

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/y3owk1n/neru/internal/derrors"
)

// ActionFeedCmd feeds one or more keys or key chords directly to the operating system.
var ActionFeedCmd = &cobra.Command{
	Use:   "feed <key> [key...]",
	Short: "Feed keys directly to the operating system",
	Long: `Feed one or more keys or key chords to the operating system or to Homearchy's
own mode system.

By default, keys are posted directly to the OS. Use --mode to route keys
through Homearchy's active mode/action pipeline instead.

Examples:
  homearchy action feed o
  homearchy action feed ctrl+c
  homearchy action feed Cmd+Shift+P
  homearchy action feed h e l o return
  homearchy action feed --mode o
  homearchy action feed --mode Escape
  homearchy action feed --mode Cmd+Shift+p`,
	Args: validateActionFeedArgs,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return requiresRunningInstance()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		toMode, _ := cmd.Flags().GetBool("mode")

		capacity := len(args) + 2 //nolint:mnd // "feed" + optional "--mode" + keys

		actionArgs := make([]string, 0, capacity)

		actionArgs = append(actionArgs, "feed")
		if toMode {
			actionArgs = append(actionArgs, "--mode")
		}

		for _, arg := range args {
			actionArgs = append(actionArgs, strings.TrimSpace(arg))
		}

		return sendCommand(cmd, "action", actionArgs)
	},
}

func init() {
	ActionFeedCmd.Flags().Bool("mode", false,
		"Feed keys into Homearchy's own mode system instead of the OS")
}

func validateActionFeedArgs(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return derrors.New(
			derrors.CodeInvalidInput,
			"feed requires at least one key (e.g., homearchy action feed o, homearchy action feed ctrl+c)",
		)
	}

	for _, arg := range args {
		if strings.TrimSpace(arg) == "" {
			return derrors.New(
				derrors.CodeInvalidInput,
				"feed keys cannot be empty",
			)
		}
	}

	return nil
}
