package xplane

// flightloop, high freq code!
func (s *xplaneService) flightLoop(
	elapsedSinceLastCall,
	elapsedTimeSinceLastFlightLoop float32,
	counter int,
	ref interface{},
) float32 {

	//if int(counter/100)%2 == 0 {
	//	s.Logger.Debugf("flightLoop: %v, %v, %v, %v", elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop, counter, ref)
	//	honeycomb.LED_STATE_CHANGED = true
	//	s.Logger.Debugf("FORCE LED_STATE_CHANGED: %v", honeycomb.LED_STATE_CHANGED)
	//}

	return -1
}
