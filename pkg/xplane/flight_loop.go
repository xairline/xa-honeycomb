package xplane

// flightloop, high freq code!
func (s *xplaneService) flightLoop(
	elapsedSinceLastCall,
	elapsedTimeSinceLastFlightLoop float32,
	counter int,
	ref interface{},
) float32 {

	if s.profile == nil {
		s.Logger.Info("Profile is nil, try to load it again")
		s.tryLoadProfile()
	}

	s.updateLeds()

	return 0.1
}
