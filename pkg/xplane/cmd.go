package xplane

import (
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/goplane/xplm/utilities"
	"github.com/xairline/xa-honeycomb/pkg"
	"time"
)

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
