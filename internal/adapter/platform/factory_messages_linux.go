//go:build linux

package platform

import (
	"os"

	"github.com/y3owk1n/neru/internal/derrors"
)

func unsupportedLinuxBackendError(backend LinuxBackend) error {
	switch backend {
	case BackendWaylandGNOME:
		return derrors.New(
			derrors.CodeNotSupported,
			"Homearchy requires an Omarchy Hyprland Wayland session; GNOME Wayland is not supported.",
		)
	case BackendWaylandOther:
		return derrors.Newf(
			derrors.CodeNotSupported,
			"Homearchy requires an Omarchy Hyprland Wayland session (XDG_CURRENT_DESKTOP=%q).",
			os.Getenv("XDG_CURRENT_DESKTOP"),
		)
	case BackendUnknown:
		return derrors.New(
			derrors.CodeNotSupported,
			"Homearchy could not detect a Wayland display server. Ensure WAYLAND_DISPLAY is set inside Omarchy.",
		)
	case BackendX11, BackendWaylandWlroots, BackendWaylandKDE:
		return derrors.Newf(
			derrors.CodeInternal,
			"unsupportedLinuxBackendError called on supported backend: %s",
			backend.String(),
		)
	default:
		return derrors.Newf(
			derrors.CodeNotSupported,
			"unsupported linux backend: %s",
			backend.String(),
		)
	}
}

func unsupportedHomearchyBackendError(backend LinuxBackend) error {
	return derrors.Newf(
		derrors.CodeNotSupported,
		"Homearchy requires an Omarchy Hyprland Wayland session; detected %s (XDG_CURRENT_DESKTOP=%q)",
		backend.String(),
		os.Getenv("XDG_CURRENT_DESKTOP"),
	)
}
