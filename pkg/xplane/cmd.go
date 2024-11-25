package xplane

import "github.com/xairline/goplane/xplm/utilities"

func (s *xplaneService) Increase(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandEnd {
		s.Logger.Debugf("Increase: %v, Phase: %v", command, phase)
	}
	return 0
}

func (s *xplaneService) Decrease(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandEnd {
		s.Logger.Debugf("Decrease: %v, Phase: %v", command, phase)
	}
	return 0
}
