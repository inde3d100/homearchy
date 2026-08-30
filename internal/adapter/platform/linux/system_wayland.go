//go:build linux

package linux

import "image"

var globalWlrootsModifierDispatcher = newWlrootsModifierDispatcher(waylandModifierEvent)

// WaylandMoveCursorToPoint moves the pointer to an absolute position.
//
// This and the other exported entry points here route to whichever injection
// backend the running compositor supports: zwlr_virtual_pointer on wlroots, or
// libei through the RemoteDesktop portal on KWin. The dispatch itself is in
// system_wayland_input.go.
func WaylandMoveCursorToPoint(point image.Point) error {
	return waylandMoveCursorToPoint(point)
}

// WaylandCursorPosition returns the cached cursor position.
func WaylandCursorPosition() (image.Point, error) {
	return waylandCursorPosition()
}

// WaylandAimVirtualPointerAtCursor warps the injection pointer to the physical
// cursor so a following axis or button event is delivered to the window under
// it. Clicks already do this; scroll has to as well — compositors ignore axis
// events from a virtual pointer that has never sent motion.
func WaylandAimVirtualPointerAtCursor() error {
	return waylandAimVirtualPointerAtCursor()
}

// WaylandClick performs a full click (press + release) at the given position.
func WaylandClick(point image.Point, button int) error {
	return waylandClick(point, button)
}

// WaylandButtonEvent presses or releases a button at the given position.
func WaylandButtonEvent(point image.Point, button int, pressed bool) error {
	return waylandButtonEvent(point, button, pressed)
}

// WaylandButtonRelease releases a button at the current cursor position.
func WaylandButtonRelease(button int) error {
	return waylandButtonRelease(button)
}

// WaylandScroll sends a scroll event. axis: 0=vertical, 1=horizontal.
// delta is in logical pixels (positive = down/right, negative = up/left).
// discrete is the discrete step count (e.g. +/-1 per logical scroll click).
// Each call emits a single scroll event; callers should loop for larger
// scroll distances.
func WaylandScroll(axis, delta, discrete int) error {
	return waylandScroll(axis, delta, discrete)
}

// WaylandScrollContinuous sends one scroll step of an arbitrary distance,
// including less than one wheel notch. axis: 0=vertical, 1=horizontal; delta is
// in the same logical pixels WaylandScroll takes, with the same sign
// convention.
//
// It exists beside WaylandScroll because the two say different things to a
// client: WaylandScroll's discrete step count declares "one wheel notch", which
// a compositor is entitled to hold back until whole notches accumulate, while
// this one declares a continuous distance and is delivered as written. Smooth
// scroll is the caller.
func WaylandScrollContinuous(axis int, delta float64) error {
	return waylandScrollContinuous(axis, delta)
}

// WlrootsScrollBatch sends multiple scroll events in a single flush.
// deltas and discretes must have the same length. Routes through the
// waylandScrollBatch seam so KDE (libei, no virtual pointer) emits one
// libeiScroll event per delta instead of taking the wlroots-only batch path
// (which fails on KWin with "failed to perform wlroots batch scroll").
func WlrootsScrollBatch(axis int, deltas, discretes []int) error {
	return waylandScrollBatch(axis, deltas, discretes)
}

// WaylandModifierEvent presses or releases a modifier key.
func WaylandModifierEvent(modifier string, isDown bool) error {
	return globalWlrootsModifierDispatcher.event(modifier, isDown)
}

// WaylandKeyEvent presses (pressed=true) or releases (pressed=false) a single
// key identified by its evdev keycode using the virtual keyboard protocol.
// This is used to inject synthetic release events for keys that were held
// before an evdev grab and released during it, so the compositor's key state
// stays consistent.
func WaylandKeyEvent(keycode uint32, pressed bool) error {
	return waylandKeyEvent(keycode, pressed)
}

// WaylandFocusedAppID returns the app_id of the focused toplevel. The
// wlr-foreign-toplevel-management protocol is the first source, because it is
// live and lock-step with the compositor. When that protocol is absent or
// silent — Hyprland 0.56 often does not bind it for ordinary clients — the
// compositor CLI (hyprctl / swaymsg / niri) answers the same question. GNOME
// has neither, so the bool stays false there. The protocol does not expose a
// PID; app_id is the identifier per-app configuration is keyed on.
func WaylandFocusedAppID() (string, bool) {
	if appID, ok := wlrootsFocusedAppID(); ok && appID != "" {
		return appID, true
	}

	appID, _, ok := compositorFocusedAppIdentity()

	return appID, ok
}

// WaylandFocusedAppIdentity returns the focused toplevel's app_id and title
// together, read so they always describe the same window. The title
// disambiguates multiple windows of the focused application, which share an
// app_id. Protocol first, compositor CLI if the protocol reports nothing; the
// bool is false when nothing is focused, GNOME/Mutter, or neither source
// answers.
func WaylandFocusedAppIdentity() (string, string, bool) {
	if appID, title, ok := wlrootsFocusedAppIdentity(); ok && appID != "" {
		return appID, title, true
	}

	return compositorFocusedAppIdentity()
}
