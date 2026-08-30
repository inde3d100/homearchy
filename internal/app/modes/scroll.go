package modes

import (
	"time"

	"go.uber.org/zap"

	"github.com/y3owk1n/neru/internal/domain"
	"github.com/y3owk1n/neru/internal/ports"
)

// StartInteractiveScroll activates the interactive scroll mode,
// showing the scroll overlay and enabling key handling for scrolling.
func (h *handlerState) startInteractiveScroll() {
	// Super+Z is bound both in Hyprland (`bin/homearchy scroll`) and in the
	// engine [hotkeys] table. The two fires land a few milliseconds apart; a
	// second enter tears the event tap down and the re-grab is ~200ms late, so
	// j/k typed in that window go to the focused app. --toggle is answered
	// before a mode is reached, so a chord that is not a toggle must be a no-op
	// when scroll is already up.
	if h.appState.CurrentMode() == domain.ModeScroll && h.scroll.Context.IsActive() {
		h.logger.Debug("Scroll mode already active")
		return
	}

	h.prepareForModeActivation()
	h.cursorState.SkipNextRestore()

	h.scroll.Context.Reset()

	if h.appState.CurrentMode() != domain.ModeIdle {
		// Mode-to-mode transition: clean up the current mode but keep the
		// event tap enabled, for the reason exitModeForTransition states — this
		// is where that dead window was first found, in the shape of missed
		// scrolling keys when activating from grid mode.
		//
		// Scroll keeps a cleanup of its own rather than calling that helper,
		// and the three differences are all deliberate: it does not clear the
		// overlay frame (the ScrollFrame below replaces it, and clearing first
		// is a wasted round trip on the backend where this mode is entered
		// most), it does not pass through idle (nothing here can abandon the
		// activation, so there is no state to be caught in), and it *does*
		// reset the suppressed modifiers that performCommonCleanup deliberately
		// keeps. Unify them only with an answer for that third one.
		h.performModeSpecificCleanup()
		h.stopHeldRepeat()

		if h.refreshHintsTimer != nil {
			h.refreshHintsTimer.Stop()
			h.refreshHintsTimer = nil
		}

		h.hotkeyLastKey = ""
		h.hotkeyLastKeyTime = 0
		h.clearStickyModifiers()
		h.releaseHeldButtons()

		h.suppressedModifiers = 0
		h.suppressedUntil = time.Time{}
		h.cursorState.Reset()

		if h.appState.HotkeyRefreshPending() {
			h.appState.SetHotkeyRefreshPending(false)

			if h.refreshHotkeys != nil {
				go h.refreshHotkeys()
			}
		}

		h.logger.Debug("Transitioned to scroll mode",
			zap.String("from", h.CurrModeString()))
	}

	h.scroll.Context.SetIsActive(true)

	h.enterMode(domain.ModeScroll)

	// Scroll draws nothing of its own, but entering it is still a transition:
	// the frame takes the previous mode's drawing off the shared surface and
	// tells the overlay which mode the indicators are naming.
	h.showFrame(ports.ScrollFrame{}, "show scroll overlay")

	h.logger.Info("Interactive scroll activated")
}

// handleGenericScrollKey intentionally does nothing.
// Scroll key behavior is fully driven by hotkeys.
func (h *handlerState) handleGenericScrollKey(_ string) {}
