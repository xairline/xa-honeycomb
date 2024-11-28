package xplane

import (
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/goplane/xplm/utilities"
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
		var myProfile profile
		var factor float64
		switch s.apSelector {
		case "ias":
			myProfile = s.profile.AP_IAS
			factor = 1
		case "alt":
			myProfile = s.profile.AP_ALT
			factor = 100
		case "vs":
			myProfile = s.profile.AP_VS
			factor = 1
		case "hdg":
			myProfile = s.profile.AP_HDG
			factor = 1
		case "crs":
			myProfile = s.profile.AP_CRS
			factor = 1
		}
		s.adjust(myProfile, direction, multiplier, factor)
		s.Logger.Infof("Knob turn: %s, Mode: %s, Multiplier: %.1f", direction, s.apSelector, multiplier)
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

func (s *xplaneService) adjust(myProfile profile, direction int, multiplier float64, factor float64) {
	if myProfile.Commands != nil {
		//TODO: Implement this
		s.Logger.Error("Not implemented")
	}
	myDataref, found := dataAccess.FindDataRef(myProfile.Datarefs[0].Dataref_str)
	if !found {
		s.Logger.Errorf("Dataref not found: %s", myProfile.Datarefs[0].Dataref_str)
		return
	}
	currentValueType := dataAccess.GetDataRefTypes(myDataref)
	switch currentValueType {
	case dataAccess.TypeFloat:
		currentValue := dataAccess.GetFloatData(myDataref)
		newValue := currentValue + float32(float64(direction)*multiplier*factor)
		s.Logger.Infof("Current Value: %f, New Value: %f", currentValue, newValue)
		dataAccess.SetFloatData(myDataref, newValue)
	case dataAccess.TypeInt:
		currentValue := dataAccess.GetIntData(myDataref)
		newValue := currentValue + int(float64(direction)*multiplier*factor)
		s.Logger.Infof("Current Value: %f, New Value: %f", currentValue, newValue)
		dataAccess.SetIntData(myDataref, newValue)
	}

}
