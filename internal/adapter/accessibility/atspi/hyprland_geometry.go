//go:build linux

package atspi

import (
	"image"
	"strings"

	"go.uber.org/zap"

	"github.com/y3owk1n/neru/internal/adapter/platform/compositorcli"
)

// Hyprland window-origin source. `hyprctl -j activewindow` reports the focused
// window's absolute position ("at") and size ("size"), which give the screen
// origin directly.
type hyprlandOriginSource struct {
	logger   *zap.Logger
	lastSize image.Point
}

func newHyprlandOriginSource(logger *zap.Logger) *hyprlandOriginSource {
	if logger == nil {
		logger = zap.NewNop()
	}

	return &hyprlandOriginSource{logger: logger.Named("accessibility.hyprland")}
}

func (h *hyprlandOriginSource) start() {}

// hyprlandWindow mirrors the fields of `hyprctl -j activewindow` we use.
// "at" is [x, y] and "size" is [w, h], both in absolute screen pixels.
type hyprlandWindow struct {
	At    []int  `json:"at"`
	Size  []int  `json:"size"`
	Class string `json:"class"`
	Title string `json:"title"`
}

func (h *hyprlandOriginSource) originFor(frame windowFrame) (image.Point, bool, error) {
	var win hyprlandWindow

	err := compositorcli.Query(&win, "hyprctl", "-j", "activewindow")
	if err != nil {
		return image.Point{}, false, err
	}

	if len(win.Size) >= coordPairLen {
		h.lastSize = image.Pt(win.Size[0], win.Size[1])
	}

	origin, ok := hyprlandComputeOrigin(win, frame, h.logger)

	return origin, ok, nil
}

func (h *hyprlandOriginSource) contentSize() image.Point {
	return h.lastSize
}

// hyprlandComputeOrigin derives the focused window's screen origin from
// `hyprctl activewindow`. A size match is the usual proof it is the same
// window as the AT-SPI frame. Chromium (and similar) report AT-SPI extents in
// CSS pixels, which can differ from Hyprland's logical size by hundreds of
// pixels; when the compositor class still matches the focused app_id, the
// origin is accepted and callers scale coordinates onto that window.
func hyprlandComputeOrigin(
	win hyprlandWindow,
	frame windowFrame,
	logger *zap.Logger,
) (image.Point, bool) {
	if len(win.At) < coordPairLen || len(win.Size) < coordPairLen {
		return image.Point{}, false
	}

	sizeOK := absInt(win.Size[0]-frame.Width) <= windowOriginSizeTolerance &&
		absInt(win.Size[1]-frame.Height) <= windowOriginSizeTolerance
	idOK := frame.FocusedAppID != "" &&
		strings.EqualFold(strings.TrimSpace(win.Class), strings.TrimSpace(frame.FocusedAppID))

	if !sizeOK && !idOK {
		logger.Debug("hyprland origin rejected: window size does not match AT-SPI frame",
			zap.Ints("windowSize", win.Size),
			zap.Int("frameW", frame.Width), zap.Int("frameH", frame.Height),
			zap.String("windowClass", win.Class),
			zap.String("focusedAppID", frame.FocusedAppID))

		return image.Point{}, false
	}

	return image.Pt(win.At[0], win.At[1]), true
}
