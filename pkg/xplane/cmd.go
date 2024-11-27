package xplane

import "github.com/xairline/goplane/xplm/utilities"

func (s *xplaneService) changeApValue(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int {
	if phase == utilities.Phase_CommandEnd {
		if ref.(string) == "up" {
			s.Logger.Infof("Increase: %v, Phase: %v, AP Mode: %s", command, phase, s.apSelector)
		} else {
			s.Logger.Infof("Decrease: %v, Phase: %v, AP Mode: %s", command, phase, s.apSelector)
		}
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
