package xplane

import "github.com/xairline/goplane/xplm/dataAccess"

// flightloop, high freq code!
func (s *xplaneService) flightLoop(
	elapsedSinceLastCall,
	elapsedTimeSinceLastFlightLoop float32,
	counter int,
	ref interface{},
) float32 {

	if s.profile == nil {
		s.Logger.Debugf("Profile is nil, try to load it again")
		aircraftIACODrf, found := dataAccess.FindDataRef("sim/aircraft/view/acf_ICAO")
		if !found {
			s.Logger.Errorf("Failed to find ICAO")
		}
		aircraftIACO := dataAccess.GetString(aircraftIACODrf)
		s.setupDataRefs(aircraftIACO)
	}

	s.updateLeds()

	return 0.1
}
