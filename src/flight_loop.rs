use xplm::flight_loop::FlightLoopCallback;

pub struct FlightLoopHandler;

impl FlightLoopCallback for FlightLoopHandler {
    fn flight_loop(&mut self, state: &mut xplm::flight_loop::LoopState) {
        // In our flight loop callback, our datarefs should be ready, and we can read the loadout
        // from file and restore it into the sim.
        return;
    }
}
