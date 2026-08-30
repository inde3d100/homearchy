# Changelog

## 0.1.0 - 2026-08-08

- Added the Omarchy service and optional bar widget.
- Bundled native `arm64` and `amd64` Homearchy engines behind an architecture dispatcher.
- Added hints, searchable hints, grid, recursive grid, scroll, and monitor-selection controls.
- Added the idempotent `homearchy-setup` command for managed Hyprland bindings and legacy Neru binding migration.
- Bound recursive grid to `Super+Shift+R` and scroll to `Super+Shift+Z` so the defaults stay off stock Omarchy Super+letter chords.
- Documented that the Super shortcuts are configurable in `bindings.lua` and `Service.qml`.
- Kept `Super+Enter` as Omarchy's terminal shortcut; Homearchy does not bind it.
- Replaced the GitHub/marketplace preview with a desktop-wallpaper screenshot.
- Bound `Super+Shift+Escape` as an emergency overlay cancel, and stopped a Wayland capture race that could freeze the session until reboot.
- Bound `Space` to left-click and `Enter` to right-click in grid and recursive grid, then return to idle.
- Inverted scroll on Omarchy so vim `j`/`k` match down/up.
- On Hyprland, read the focused window from `hyprctl` when the Wayland foreign-toplevel protocol does not report one, so hints can label the active app.
- On Chromium-style apps that expose extra untitled AT-SPI frames, pick the unique named or active window of the focused app so hints still appear.
- Scale Chromium AT-SPI coordinates onto Hyprland's logical window size so hint labels are not filtered off-screen.
- When a window exposes no clickable accessibility targets (Chromium on Wayland), fall back to screen detection so hints still appear.
- Added Omarchy theme synchronization, engine health polling, supervised restarts, SHA-256 checksums, and build-provenance attestation.
