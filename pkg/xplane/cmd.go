package xplane

import "C"
import (
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/goplane/xplm/utilities"
	"github.com/xairline/xa-honeycomb/pkg"
	"time"
)

const doubleClickThreshold = 500 * time.Millisecond // Define double-click threshold

func (s *xplaneService) changeApValue(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	// Handle only when command phase is CommandEnd
	if phase == utilities.Phase_CommandEnd {
		now := time.Now()

		// Determine speed multiplier based on time elapsed
		var multiplier float64
		if !s.lastKnobTime.IsZero() {
			elapsed := now.Sub(s.lastKnobTime).Milliseconds()
			if elapsed < 100 {
				multiplier = 5.0 // Fast turn
			} else if elapsed < 200 {
				multiplier = 3.0 // Medium turn
			} else {
				multiplier = 1.0 // Slow turn
			}
		} else {
			multiplier = 1.0
		}

		direction := 0
		// Log the adjustment
		if ref.(string) == "up" {
			s.Logger.Debugf("Increase: %v, Phase: %v, AP Mode: %s, Multiplier: %.1f", command, phase, s.apSelector, multiplier)
			direction = 1
		} else {
			s.Logger.Debugf("Decrease: %v, Phase: %v, AP Mode: %s, Multiplier: %.1f", command, phase, s.apSelector, multiplier)
			direction = -1
		}
		var myProfile pkg.KnobProfile
		var step float64
		switch s.apSelector {
		case "ias":
			myProfile = s.profile.Knobs.AP_IAS

			iasStep, foundIasStep := s.dataValue(&s.profile.Data.AP_IAS_STEP)
			if foundIasStep {
				step = iasStep
			} else {
				step = 1
			}
		case "alt":
			myProfile = s.profile.Knobs.AP_ALT

			altStep, foundAltStep := s.dataValue(&s.profile.Data.AP_ALT_STEP)
			if foundAltStep {
				step = altStep
			} else {
				step = 100
			}
		case "vs":
			myProfile = s.profile.Knobs.AP_VS

			vsStep, foundVsStep := s.dataValue(&s.profile.Data.AP_VS_STEP)
			if foundVsStep {
				step = vsStep
			} else {
				step = 1
			}
		case "hdg":
			myProfile = s.profile.Knobs.AP_HDG
			step = 1
		case "crs":
			myProfile = s.profile.Knobs.AP_CRS
			step = 1
		}
		s.adjust(myProfile, direction, multiplier, step)
		s.Logger.Debugf("Knob turn: %d, Mode: %s, Multiplier: %.1f, Step: %.1f", direction, s.apSelector, multiplier, step)
		// Update the last interaction time
		s.lastKnobTime = now
	}
	return 0
}

func (s *xplaneService) changeAPMode(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if s.apSelector != ref.(string) {
		s.Logger.Debugf("AP MODE CHANGE: %v, Phase: %v, ref: %s", command, phase, ref.(string))
		s.apSelector = ref.(string)
	}
	return 0
}

func (s *xplaneService) adjust(myProfile pkg.KnobProfile, direction int, multiplier float64, step float64) {
	if myProfile.Commands != nil {
		var cmd utilities.CommandRef
		if direction > 0 {
			cmd = utilities.FindCommand(myProfile.Commands[0].CommandStr)
		} else {
			cmd = utilities.FindCommand(myProfile.Commands[1].CommandStr)
		}
		utilities.CommandOnce(cmd)
	}

	for i := 0; i < len(myProfile.Datarefs); i++ {
		myDatarefName := myProfile.Datarefs[i].DatarefStr
		myDataref, found := dataAccess.FindDataRef(myDatarefName)
		if !found {
			s.Logger.Errorf("Dataref[%d] not found: %s", i, myDatarefName)
			continue
		}
		currentValueType := dataAccess.GetDataRefTypes(myDataref)
		switch currentValueType {
		case dataAccess.TypeFloat:
			currentValue := dataAccess.GetFloatData(myDataref)
			newValue := currentValue + float32(float64(direction)*multiplier*step)
			s.Logger.Debugf("Knob dataref: %s, Current Value: %f, New Value: %f", myDatarefName, currentValue, newValue)
			dataAccess.SetFloatData(myDataref, newValue)
		case dataAccess.TypeInt:
			currentValue := dataAccess.GetIntData(myDataref)
			newValue := currentValue + int(float64(direction)*multiplier*step)
			s.Logger.Debugf("Knob dataref: %s, Current Value: %f, New Value: %f", myDatarefName, currentValue, newValue)
			dataAccess.SetIntData(myDataref, newValue)
		}
	}
}

func (s *xplaneService) setupKnobsCmds() {
	increaseCmd := utilities.CreateCommand("Honeycomb Bravo/increase", "Increase the value of the autopilot mode selected with the rotary encoder.")
	decreaseCmd := utilities.CreateCommand("Honeycomb Bravo/decrease", "Decrease the value of the autopilot mode selected with the rotary encoder.")

	mode_ias := utilities.CreateCommand("Honeycomb Bravo/mode_ias", "Set the autopilot mode to IAS.")
	mode_alt := utilities.CreateCommand("Honeycomb Bravo/mode_alt", "Set the autopilot mode to ALT.")
	mode_vs := utilities.CreateCommand("Honeycomb Bravo/mode_vs", "Set the autopilot mode to VS.")
	mode_hdg := utilities.CreateCommand("Honeycomb Bravo/mode_hdg", "Set the autopilot mode to HDG.")
	mode_crs := utilities.CreateCommand("Honeycomb Bravo/mode_crs", "Set the autopilot mode to CRS.")

	// set up command handlers
	utilities.RegisterCommandHandler(increaseCmd, s.changeApValue, true, "up")
	utilities.RegisterCommandHandler(decreaseCmd, s.changeApValue, true, "down")
	utilities.RegisterCommandHandler(mode_ias, s.changeAPMode, true, "ias")
	utilities.RegisterCommandHandler(mode_alt, s.changeAPMode, true, "alt")
	utilities.RegisterCommandHandler(mode_vs, s.changeAPMode, true, "vs")
	utilities.RegisterCommandHandler(mode_hdg, s.changeAPMode, true, "hdg")
	utilities.RegisterCommandHandler(mode_crs, s.changeAPMode, true, "crs")
}

func (s *xplaneService) setupApCmds() {

	ap_ias := utilities.CreateCommand("Honeycomb Bravo/ap_ias", "Bravo IAS pressed.")
	ap_alt := utilities.CreateCommand("Honeycomb Bravo/ap_alt", "Bravo ALT pressed.")
	ap_vs := utilities.CreateCommand("Honeycomb Bravo/ap_vs", "Bravo VS pressed.")
	ap_hdg := utilities.CreateCommand("Honeycomb Bravo/ap_hdg", "Bravo HDG pressed.")
	ap_rev := utilities.CreateCommand("Honeycomb Bravo/ap_rev", "Bravo REV pressed.")
	ap_nav := utilities.CreateCommand("Honeycomb Bravo/ap_nav", "Bravo NAV pressed.")
	ap_apr := utilities.CreateCommand("Honeycomb Bravo/ap_apr", "Bravo APR pressed.")
	ap := utilities.CreateCommand("Honeycomb Bravo/ap", "Bravo AP pressed.")

	// set up command handlers
	utilities.RegisterCommandHandler(ap_ias, s.apPressed, true, "ias")
	utilities.RegisterCommandHandler(ap_alt, s.apPressed, true, "alt")
	utilities.RegisterCommandHandler(ap_vs, s.apPressed, true, "vs")
	utilities.RegisterCommandHandler(ap_hdg, s.apPressed, true, "hdg")
	utilities.RegisterCommandHandler(ap_rev, s.apPressed, true, "rev")
	utilities.RegisterCommandHandler(ap_nav, s.apPressed, true, "nav")
	utilities.RegisterCommandHandler(ap_apr, s.apPressed, true, "apr")
	utilities.RegisterCommandHandler(ap, s.apPressed, true, "ap")
}

func (s *xplaneService) apPressed(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandEnd {
		buttonRef := ref.(string) // Convert ref to string (or your button identifier type)
		now := time.Now()

		s.mutex.Lock()
		defer s.mutex.Unlock()

		// Check if there's an existing timer for the button
		if timer, exists := s.clickTimers[buttonRef]; exists {
			// Double-click detected, cancel the timer and trigger double-click logic
			timer.Stop()
			delete(s.clickTimers, buttonRef)
			s.doubleClick(buttonRef)
			return 0
		}

		// Single-click detected; set a timer to delay action
		timer := time.AfterFunc(doubleClickThreshold, func() {
			s.mutex.Lock()
			defer s.mutex.Unlock()

			// Ensure the timer is not removed by a double-click
			if s.clickTimers[buttonRef] != nil {
				delete(s.clickTimers, buttonRef)
				s.Logger.Debugf("Single-click detected for button: %s, timestamp: %s", buttonRef, now)
				s.singleClick(buttonRef)
			}
		})

		// Store the timer in the map
		s.clickTimers[buttonRef] = timer
		s.lastClickTime[buttonRef] = now
	}

	return 0
}

func (s *xplaneService) singleClick(ref string) {
	s.cmdEventQueueMu.Lock()
	defer s.cmdEventQueueMu.Unlock()
	switch ref {
	case "hdg":
		s.commands(s.profile.Buttons.HDG.SingleClick)
	case "nav":
		s.commands(s.profile.Buttons.NAV.SingleClick)
	case "alt":
		s.commands(s.profile.Buttons.ALT.SingleClick)
	case "apr":
		s.commands(s.profile.Buttons.APR.SingleClick)
	case "vs":
		s.commands(s.profile.Buttons.VS.SingleClick)
	case "ap":
		s.commands(s.profile.Buttons.AP.SingleClick)
	case "rev":
		s.commands(s.profile.Buttons.REV.SingleClick)
	case "ias":
		s.commands(s.profile.Buttons.IAS.SingleClick)
	default:
		s.Logger.Debugf("Single-click detected for button: %s", ref)
	}
}

func (s *xplaneService) doubleClick(ref string) {
	s.cmdEventQueueMu.Lock()
	defer s.cmdEventQueueMu.Unlock()
	switch ref {
	case "hdg":
		s.commands(s.profile.Buttons.HDG.DoubleClick)
	case "nav":
		s.commands(s.profile.Buttons.NAV.DoubleClick)
	case "alt":
		s.commands(s.profile.Buttons.ALT.DoubleClick)
	case "apr":
		s.commands(s.profile.Buttons.APR.DoubleClick)
	case "vs":
		s.commands(s.profile.Buttons.VS.DoubleClick)
	case "ap":
		s.commands(s.profile.Buttons.AP.DoubleClick)
	case "rev":
		s.commands(s.profile.Buttons.REV.DoubleClick)
	case "ias":
		s.commands(s.profile.Buttons.IAS.DoubleClick)
	default:
		s.Logger.Debugf("Double-click detected for button: %s", ref)
	}
}
func (s *xplaneService) commands(commands []pkg.Command) {
	for _, cmd := range commands {
		s.cmdEventQueue = append(s.cmdEventQueue, cmd.CommandStr)
	}
}
