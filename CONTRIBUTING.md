# Contributing to Homearchy

Homearchy targets Omarchy 4.0 or newer on Hyprland/Wayland. Changes must preserve the bundled `amd64` and `arm64` engine contract and must not add a dependency on an installed Neru service or binary.

## Local checks

Install the Linux development libraries listed in [README.md](README.md), then run:

```bash
go test ./...
omarchy plugin validate .
(cd bin && sha256sum -c SHA256SUMS)
hyprctl reload
hyprctl configerrors
```

Run `go test` once after all source changes; this repository has native CGO packages and a full run is intentionally expensive. Never run desktop-driving tests while Neru or another Homearchy daemon is active.

For plugin work, place the checkout at `~/.config/omarchy/plugins/io.github.inde3d100.homearchy` or link that path during development, enable it with `omarchy plugin enable io.github.inde3d100.homearchy right`, then rescan with:

```bash
omarchy-shell shell rescanPlugins
```

Verify service state with:

```bash
omarchy-shell io.github.inde3d100.homearchy status
```

## Pull requests

Keep changes scoped. Update observable behavior, its focused tests, the manifest or README when applicable, and `CHANGELOG.md` for user-facing changes. Do not commit generated caches, local state, or secrets.

Native engine releases are built on native GitHub runners through `.github/workflows/engine-build.yml`. Do not replace either tracked architecture binary with a placeholder or a cross-architecture wrapper.
