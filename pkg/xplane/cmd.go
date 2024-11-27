package xplane

import "github.com/xairline/goplane/xplm/utilities"

func (s *xplaneService) Increase(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandEnd {
		s.Logger.Infof("Increase: %v, Phase: %v", command, phase)
		if s.apSelector == "" {
			s.Logger.Infof("AP MODE NOT SET")
			return 0
		}
	}
	return 0
}

func (s *xplaneService) Decrease(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandEnd {
		s.Logger.Infof("Decrease: %v, Phase: %v", command, phase)
		if s.apSelector == "" {
			s.Logger.Infof("AP MODE NOT SET")
			return 0
		}
	}
	return 0
}

func (s *xplaneService) changeAPMode(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandBegin || (phase == utilities.Phase_CommandContinue && s.apSelector != ref.(string)) {
		s.Logger.Infof("AP MODE CHANGE: %v, Phase: %v, ref: %s", command, phase, ref.(string))
		s.apSelector = ref.(string)
	}
	return 0
}
