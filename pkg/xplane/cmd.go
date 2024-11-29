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
		var myProfile pkg.BravoProfile
		var step float64
		switch s.apSelector {
		case "ias":
			myProfile = s.profile.Knobs.AP_IAS
			step = 1
		case "alt":
			myProfile = s.profile.Knobs.AP_ALT

			altStep := s.valueOf(&s.profile.Data.AP_ALT_STEP)
			if altStep != 0 {
				step = altStep
			} else {
				step = 100
			}
		case "vs":
			myProfile = s.profile.Knobs.AP_VS

			vsStep := s.valueOf(&s.profile.Data.AP_VS_STEP)
			if vsStep != 0 {
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
		s.Logger.Infof("Knob turn: %d, Mode: %s, Multiplier: %.1f, Step: %.1f", direction, s.apSelector, multiplier, step)
		// Update the last interaction time
		s.lastKnobTime = now
	}
	return 0
}

func (s *xplaneService) changeAPMode(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if s.apSelector != ref.(string) {
		s.Logger.Infof("AP MODE CHANGE: %v, Phase: %v, ref: %s", command, phase, ref.(string))
		s.apSelector = ref.(string)
	}
	return 0
}

func (s *xplaneService) adjust(myProfile pkg.BravoProfile, direction int, multiplier float64, step float64) {
	if myProfile.Commands != nil {
		var cmd utilities.CommandRef
		if direction > 0 {
			cmd = utilities.FindCommand(myProfile.Commands[0].CommandStr)
		} else {
			cmd = utilities.FindCommand(myProfile.Commands[1].CommandStr)
		}
		utilities.CommandOnce(cmd)
	}
	myDataref, found := dataAccess.FindDataRef(myProfile.Datarefs[0].DatarefStr)
	if !found {
		s.Logger.Errorf("Dataref not found: %s", myProfile.Datarefs[0].DatarefStr)
		return
	}
	currentValueType := dataAccess.GetDataRefTypes(myDataref)
	switch currentValueType {
	case dataAccess.TypeFloat:
		currentValue := dataAccess.GetFloatData(myDataref)
		newValue := currentValue + float32(float64(direction)*multiplier*step)
		s.Logger.Infof("Current Value: %f, New Value: %f", currentValue, newValue)
		dataAccess.SetFloatData(myDataref, newValue)
	case dataAccess.TypeInt:
		currentValue := dataAccess.GetIntData(myDataref)
		newValue := currentValue + int(float64(direction)*multiplier*step)
		s.Logger.Infof("Current Value: %f, New Value: %f", currentValue, newValue)
		dataAccess.SetIntData(myDataref, newValue)
	}

}
