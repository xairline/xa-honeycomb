package xplane

import "github.com/xairline/xa-honeycomb/pkg/honeycomb"

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
		s.lastCounter = 0
	}

	if counter-s.lastCounter > 200 {
		honeycomb.LED_STATE_CHANGED_LOCK.Lock()
		honeycomb.LED_STATE_CHANGED = true
		honeycomb.LED_STATE_CHANGED_LOCK.Unlock()
		s.lastCounter = counter
		s.Logger.Infof("force led sync, counter: %d", counter)
	}

	s.updateLeds()

	return 0.1
}
