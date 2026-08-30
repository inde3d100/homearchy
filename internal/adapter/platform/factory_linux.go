//go:build linux

package platform

import (
	"context"
	"fmt"
	"os"

	"github.com/y3owk1n/neru/internal/adapter/platform/linux"
	"github.com/y3owk1n/neru/internal/ports"
)

// NewSystemPort returns a Linux SystemPort implementation.
func NewSystemPort() (ports.SystemPort, error) {
	backend := DetectLinuxBackend()
	if backend == BackendWaylandWlroots && isHyprlandSession() {
		return linux.NewSystemAdapter(backend.String()), nil
	}

	switch backend {
	case BackendUnknown, BackendWaylandGNOME, BackendWaylandOther:
		return nil, unsupportedLinuxBackendError(backend)
	case BackendX11, BackendWaylandWlroots, BackendWaylandKDE:
		return nil, unsupportedHomearchyBackendError(backend)
	default:
		return nil, unsupportedLinuxBackendError(backend)
	}
}

// NewFontResolver returns a Linux-backed FontResolver backed by fontconfig
// (CGO builds) or a no-CGO passthrough that still maps generic aliases.
func NewFontResolver() ports.FontResolver {
	return linux.NewFontResolver()
}

// ShowConfigOnboardingAlert tells a first-time user that Homearchy started on
// built-in defaults, and how to get a config file, then answers with that
// choice.
//
// Linux answers instead of asking. A modal dialog would add a toolkit
// dependency or rely on a helper that a minimal Omarchy session need not have.
// Starting on safe defaults keeps plugin startup non-blocking.
func ShowConfigOnboardingAlert(configPath string) int {
	title := "Homearchy is running on built-in defaults"
	message := "No configuration file at " + configPath +
		". Run `homearchy config init` to create one."

	// Bounded by the adapter's own notify deadline; these run before the daemon
	// is up, so a wedged session bus costs a moment rather than the launch.
	err := linux.ShowAlert(context.Background(), title, message)
	if err != nil {
		// Onboarding has no other channel: nothing upstream prints this, so a
		// session with no notification daemon would otherwise learn nothing.
		fmt.Fprintf(os.Stderr, "⚠️  %s.\n%s\n\n", title, message)
	}

	return ConfigOnboardingDefaults
}

// ShowConfigValidationErrorAlert puts a rejected configuration in front of the
// user before Homearchy exits.
func ShowConfigValidationErrorAlert(errorMessage, configPath string) int {
	// The launcher has already written the same failure to stderr, so a
	// missing notification daemon costs the desktop copy rather than the
	// message — nothing to fall back to here.
	_ = linux.ShowAlert(
		context.Background(),
		"Homearchy could not load "+configPath,
		errorMessage,
	)
	return ConfigValidationOK
}

// CheckAccessibilityPermissions is always true on Linux for startup gating.
func CheckAccessibilityPermissions() bool {
	return true
}

// ShowAccessibilityPermissionStartupAlert is a no-op on Linux.
func ShowAccessibilityPermissionStartupAlert() int {
	return AccessibilityPermissionStartupGranted
}
