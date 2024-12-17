// src/lib.rs
extern crate xplm;
mod bravo;
mod flight_loop;
mod misc;
pub mod profile;
pub use profile::types::Profile;
use std::os::raw::c_void;

use xplm::flight_loop::FlightLoop;
use xplm::plugin::messages::XPLM_MSG_PLANE_LOADED;
use xplm::plugin::{Plugin, PluginInfo};
use xplm::xplane_plugin;

struct XaHoneycombPlugin {
    _flight_loop: FlightLoop,
    _current_profile: Option<Profile>,
}

impl Plugin for XaHoneycombPlugin {
    type Error = std::convert::Infallible;

    fn start() -> Result<Self, Self::Error> {
        plugin_debugln!("Plugin Started");
        Ok(XaHoneycombPlugin {
            _flight_loop: FlightLoop::new(flight_loop::FlightLoopHandler),
            _current_profile: None,
        })
    }

    fn enable(&mut self) -> Result<(), Self::Error> {
        plugin_debugln!("enabling flight loop callback");
        self._flight_loop.schedule_immediate();
        Ok(())
    }

    fn disable(&mut self) {
        plugin_debugln!("disabling");
        self._flight_loop.deactivate();
    }

    fn info(&self) -> PluginInfo {
        PluginInfo {
            name: String::from("XA Honeycomb Bravo"),
            signature: String::from("com.github.xairline.xa-honeycomb"),
            description: String::from("A plugin to configure honeycomb bravo"),
        }
    }

    fn receive_message(&mut self, from: i32, message: i32, param: *mut c_void) {
        if message == XPLM_MSG_PLANE_LOADED {
            plugin_debugln!("Plane loaded");
            self._current_profile = None;
            bravo::led::all_off();
        }
    }
}

xplane_plugin!(XaHoneycombPlugin);
