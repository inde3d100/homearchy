#!/usr/bin/env bash
set -euo pipefail

export LC_ALL=C
export TZ=UTC
export CGO_ENABLED=1
export CGO_CFLAGS_ALLOW='-fno-strict-overflow'

readonly target_arch="${1:-$(go env GOARCH)}"
readonly host_arch="$(go env GOARCH)"

case "$target_arch" in
  amd64 | arm64) ;;
  *)
    printf 'Unsupported Linux architecture: %s\n' "$target_arch" >&2
    exit 64
    ;;
esac

if [[ "$target_arch" != "$host_arch" ]]; then
  printf 'CGO engine builds must run natively: host=%s target=%s\n' "$host_arch" "$target_arch" >&2
  exit 64
fi

readonly version="${VERSION:-dev}"
readonly git_commit="${GIT_COMMIT:-unknown}"
readonly build_date="${BUILD_DATE:-unknown}"
readonly output="${OUTPUT:-bin/homearchy-linux-$target_arch}"
readonly ldflags="-s -w -buildid= -X github.com/y3owk1n/neru/internal/buildinfo.Version=$version -X github.com/y3owk1n/neru/internal/buildinfo.GitCommit=$git_commit -X github.com/y3owk1n/neru/internal/buildinfo.BuildDate=$build_date"

mkdir -p "$(dirname -- "$output")"
GOOS=linux GOARCH="$target_arch" go build \
  -buildvcs=false \
  -mod=readonly \
  -trimpath \
  -ldflags="$ldflags" \
  -o "$output" \
  ./cmd/neru
