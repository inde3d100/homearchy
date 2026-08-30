# Security policy

## Reporting a vulnerability

Use [GitHub private vulnerability reporting](https://github.com/inde3d100/homearchy/security/advisories/new). Do not open a public issue for an unpatched vulnerability.

Include the Homearchy version, Omarchy version, architecture, reproduction steps, impact, and any relevant logs with secrets removed.

## Security model

Homearchy can inspect accessibility metadata and inject keyboard and pointer input. These capabilities are required for navigation and make the engine sensitive desktop-session software.

Homearchy:

- runs as the current user;
- scopes its Unix socket to the current user;
- stores state under the current user's XDG directories;
- does not include telemetry, accounts, or crash reporting;
- does not intentionally send UI content or keystrokes over the network;
- disables its own global launcher hotkeys in the plugin-generated config and uses explicit Hyprland bindings instead.

The Omarchy service starts only the architecture-matched bundled engine and supervises restarts. Release workflows publish SHA-256 checksums and GitHub build-provenance attestations for both engine binaries.

## Supported versions

Security fixes target the latest Homearchy release on Omarchy 4.0 or newer, Hyprland, and Wayland. `amd64` and `arm64` are supported.
