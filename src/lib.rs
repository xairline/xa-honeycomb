// src/lib.rs
extern crate xplm;
mod bravo;
mod misc;
mod profile;

use xplm::flight_loop::{FlightLoop, LoopState};
// Declare the `misc` module
use crate::misc::path::get_system_path;
use xplm::plugin::{Plugin, PluginInfo};
use xplm::xplane_plugin;

struct MinimalPlugin {
    _flight_loop: FlightLoop,
}

impl Plugin for MinimalPlugin {
    type Error = std::convert::Infallible;

    fn start() -> Result<Self, Self::Error> {
        plugin_debugln!("Plugin Started");
        // let xp_perf_path = get_system_path();
        // plugin_debugln!("Plugin path: {:?}", xp_perf_path);
        // match read_xplane_preferences(&xp_perf_path) {
        //     Ok(true) => {
        //         plugin_debugln!("Load flight on start is enabled.");
        //         plugin_debugln!("Mount AO and wait for all mounts to finish.");
        //     }
        //     Ok(false) => {
        //         plugin_debugln!("Load flight on start is not enabled.");
        //         plugin_debugln!("Mount AO and return immediately.");
        //     }
        //     Err(err) => eprintln!("An error occurred: {}", err),
        // }
        // xplm::flight_loop::FlightLoopCallback::new(|_, _, _, _| {
        //     plugin_debugln!("Flight loop callback");
        //     xplm::flight_loop::FlightLoopPhase::AfterFlightModel
        // });
        let mut flight_loop = FlightLoop::new(move |loop_state: &mut LoopState| {
            plugin_debugln!("Flight loop callback");
        });
        flight_loop.schedule_immediate();
        Ok(MinimalPlugin {
            _flight_loop: flight_loop,
        })
    }

    fn info(&self) -> PluginInfo {
        PluginInfo {
            name: String::from("XA Honeycomb Bravo"),
            signature: String::from("com.github.xairline.xa-honeycomb"),
            description: String::from("A plugin to configure honeycomb bravo"),
        }
    }
}

xplane_plugin!(MinimalPlugin);
