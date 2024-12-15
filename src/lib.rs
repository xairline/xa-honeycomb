// src/lib.rs
extern crate xplm;
mod logger;

use xplm::plugin::{Plugin, PluginInfo};
use xplm::xplane_plugin;

struct MinimalPlugin;

impl Plugin for MinimalPlugin {
    type Error = std::convert::Infallible;

    fn start() -> Result<Self, Self::Error> {
        // output the path of current plugin
        // let xp_perf_path = get_system_path();
        // plugin_debugln!("Plugin path: {:?}", xp_perf_path);
        // match read_xplane_preferences(&xp_perf_path) {
        //     Ok(true) => {
        //         plugin_debugln!("Load flight on start is enabled.");
        //         plugin_debugln!("Mount AO and wait for all mounts to finish.");
        //         mount(true)
        //     }
        //     Ok(false) => {
        //         plugin_debugln!("Load flight on start is not enabled.");
        //         plugin_debugln!("Mount AO and return immediately.");
        //         mount(false)
        //     }
        //     Err(err) => eprintln!("An error occurred: {}", err),
        // }

        Ok(MinimalPlugin)
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