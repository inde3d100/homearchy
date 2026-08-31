# Changelog

## 0.1.0 - 2026-08-08

- Added the Omarchy service and optional bar widget.
- Bundled native `arm64` and `amd64` Homearchy engines behind an architecture dispatcher.
- Added hints, searchable hints, grid, recursive grid, scroll, and monitor-selection controls. Monitor selection is alpha: it has not been exercised on a multi-monitor Hyprland session, and a single display is a no-op.
- Added the idempotent `homearchy-setup` command for managed Hyprland bindings and legacy Neru binding migration.
- Bound recursive grid to `Super+Shift+R` and scroll to `Super+Shift+Z` so the defaults stay off stock Omarchy Super+letter chords.
- Documented that the Super shortcuts are configurable in `bindings.lua` and `Service.qml`.
- Kept `Super+Enter` as Omarchy's terminal shortcut; Homearchy does not bind it.
- Replaced the GitHub/marketplace preview with a collage of the bar menu, full grid, and recursive grid.
- Added README screenshots of the bar menu, full grid, and recursive grid.
- Added a Configure hotkeys page to the bar menu.
- Bound `Super+Shift+Escape` as an emergency overlay cancel, and stopped a Wayland capture race that could freeze the session until reboot.
- Bound `Space` to left-click and `Enter` to right-click in grid and recursive grid, then return to idle.
- Inverted scroll on Omarchy so vim `j`/`k` match down/up.
- From scroll, the grid and recursive-grid shortcuts switch modes without Escape first. Repo default chords stay Super+Shift.
- On Hyprland, scroll at the window under the cursor without a prior trackpad move: the virtual pointer now forces a hit-test, scroll mode does not map the overlay layer, and a second Super+Z does not tear down the keyboard grab.
- On Hyprland, read the focused window from `hyprctl` when the Wayland foreign-toplevel protocol does not report one, so hints can label the active app.
- On Chromium-style apps that expose extra untitled AT-SPI frames, pick the unique named or active window of the focused app so hints still appear.
- Scale Chromium AT-SPI coordinates onto Hyprland's logical window size so hint labels are not filtered off-screen.
- When a window exposes no clickable accessibility targets (Chromium on Wayland), fall back to screen detection so hints still appear.
- Added Omarchy theme synchronization, engine health polling, supervised restarts, SHA-256 checksums, and build-provenance attestation.
